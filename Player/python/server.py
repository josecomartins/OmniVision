from concurrent.futures import ThreadPoolExecutor
import grpc
import numpy as np
import PySpin
import EasyPySpin
import cv2 as cv
print("Importing YOLO...")
from ultralytics import YOLO
print("YOLO Imported!")
import cv2
import time
from message_pb2 import Response_Omni, Object_Omni
from message_pb2_grpc import Yolo_OmniServicer, add_Yolo_OmniServicer_to_server
import time
import bisect
import sys

# Load model
model = YOLO("trained_weights/best.pt")

# Define the coordinates to find the closest point to
target = np.array([480, 480])

# Pre-compute the distance matrix
x, y = np.meshgrid(np.arange(480), np.arange(480))
distances = np.sqrt((x - target[0])**2 + (y - target[1])**2)

def sort_omni(item):
    return item.dist
    
actual_frame = np.zeros((480,480,3), dtype=np.uint8)
def open_frame_source():
    global omnicap
    try:
        omnicap = EasyPySpin.VideoCapture(0)
        omnicap.cam.PixelFormat.SetValue(PySpin.PixelFormat_RGB8Packed)
        return
    except:
        print("Can't Open Omni Camera... Verify Connection")
        omnicap = cv2.VideoCapture("MyVideo.avi")
        omnicap.set(cv.CAP_PROP_FPS,300)

last_frame = np.zeros((480,480,1), dtype=np.uint8)
actual_frame_omni = np.zeros((480,480,1), dtype=np.uint8)
def get_omni_frame(frame_bckup=1):
    global omnicap
    global model
    global distances
    global last_frame
    ret,image_test=omnicap.read()
    if(ret):
        frame_omni = cv.cvtColor(image_test, cv.COLOR_BGR2RGB)
        last_frame = frame_omni
    else:
        frame_omni= last_frame
    return 1, frame_omni
classes2 = ["Ball","Blue_Shirt","Goal", "Person","Red_Shirt","Robot"]

def Yolo_Ifer_Omni():
    global actual_frame_omni
    _, frame = get_omni_frame()
    frame_ori = frame.copy()
    results = model.predict(source=frame,conf=0.60, verbose=False)
    objects = []
    idx=0
    for index , box in enumerate(results[0].boxes):
        idx=index
        points = results[0].masks.xy[index] 
        # Find the index of the point with the smallest distance
        if len(points[:,1])>0:
            closest_index = np.argmin(distances[points[:,1].astype(int), points[:,0].astype(int)]).astype(int)
            object_atual = Object_Omni(id=int(box.cls[0]),x=int(points[closest_index][0]),y=int(points[closest_index][1]),dist=int(distances[int(points[closest_index][0])][int(points[closest_index][1])]),conf=int(box.conf[0]*100))
    if(idx>0):
        bisect.insort(objects, object_atual, key=sort_omni)
    actual_frame_omni=frame
    return frame_ori, objects
classes = ["Ball","Blue_Shirt","Goal", "Person","Red_Shirt","Robot"]

omniimg_to_send = np.zeros((480,480,1), dtype=np.uint8)

class Yolo_grpcServer_omni(Yolo_OmniServicer):
    def Send_Omni(self, request, context):
         global omniimg_to_send
         global start_time
         start=time.time()
         while 1:
            img_omni,objects = Yolo_Ifer_Omni()
            if(request.check==True):
                omniimg_to_send =np.frombuffer(request.image,dtype=np.uint8)
            resp = Response_Omni(omni = bytes(img_omni),objects = objects)
            yield resp


if __name__ == '__main__':
    global n_robot
    print("Opening OmniCamera...")
    open_frame_source()
    print("OmniCamera opened!")
    print("Opening gRPC Server...")
    server = grpc.server(ThreadPoolExecutor())
    add_Yolo_OmniServicer_to_server(Yolo_grpcServer_omni(), server)
    print("Omni Added")
    n_robot = int(sys.argv[1])
    ip_grpc = sys.argv[2]
    port = 40000
    server.add_insecure_port(f'{ip_grpc}:{port}')
    server.start()
    print(f'gRPC Server ready on IP {ip_grpc}:{port}')
    server.wait_for_termination()
    
