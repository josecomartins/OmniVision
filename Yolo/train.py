import os
HOME = os.getcwd()
print(HOME)
import ultralytics
ultralytics.checks()
from ultralytics import YOLO
from IPython.display import display, Image

## Download Dataset
from roboflow import Roboflow
rf = Roboflow(api_key="vwFRi233QozkicYHpFpz")
project = rf.workspace("omnidatasetsegmentation").project("msl-omnidirecional-lar-uminho")
dataset = project.version(14).download("yolov8")

from IPython import display
display.clear_output()
import torch
torch.cuda.empty_cache()

## Model
model = YOLO("yolov8s-seg.pt")
## Train
# On console: yolo task=segment mode=train model=yolov8s-seg.pt data=./MSL-Omnidirecional-LAR-UMinho-14/data.yaml epochs=1000 imgsz=480 plots=True
model.train(model="yolov8s-seg.pt",task="segment",data=HOME+"/MSL-Omnidirecional-LAR-UMinho-14/data.yaml", epochs=1000,imgsz=480,plots=True)
## Deploy
project.version(dataset.version).deploy(model_type="yolov8", model_path=f"/home/robot3/Desktop/Tese SLender/Testes_dataset/MSL Omnidirecional LAR UMinho.v14i.yolov8/runs1/segment/train")


