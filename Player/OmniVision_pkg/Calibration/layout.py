#!/usr/bin/env python3
"""! @brief 2Calibration of the colors of the field."""
##
# @mainpage Calibration
#
# @section description_main Description
# To easy calibrate all sensors.
# All the senttings should be save in a .txt file.
# 
# @section notes_main Calibrations
# calib_cam_omni.py -> Calibrates Omni Camera Colour.
#
# calib_cam2.py -> Calibrates Front Camera Colour .
#
# Copyright (c) 2022 MinhoTeam.  All rights reserved.
##
# @file calib_cam_omni.py
#
# @brief Calibrates Omni Camera Colour. 
#
# @section description_doxygen_example Description
# Just move the track bars to right position to see the correct colours (like: field -> green , lines -> white, robots -> black). The wanted result is a 100% of image full of vivid colours.
#
# @section libraries_main Libraries/Modules
# - time standard library (https://docs.python.org/3/library/time.html)
#   - Access to sleep function.
# - openCV standart library (https://docs.python.org/3/library/opencv.html)
#   - Access to Camera and Computer Vision function.
#
# @section notes_doxygen_example Notes
# - Comments are Doxygen compatible.
#
# \todo
# - Make Flycamera capture
# - Make Track-bars
# - Make Filter
# - Read .txt
# - Write .txt
# 
# \test 
# - Capture image from camera
#
# \bug 
# - Flycamera capture

# @section author_doxygen_example Author(s)
# - Created by José Martins on 24/09/2022.
# - Modified by José Martins on 25/09/2022.
#
# Copyright (c) 2022 MinhoTeam.  All rights reserved.

import numpy as np
import PySimpleGUI as sg

# Global Constants
## The mode of capture; 0 = camera, 1 = video, 2 = image.
capture = 1 
dists_calibration=[15,30,40,50,60,70,80,110,120,150,180,210,240,270,320,370,420,480,540,640,750,900,1000]
dists_calibration = np.array(dists_calibration)+24
label_ball_calibration=[15,30,50,80,100,150,180,240,320,420,540,640,750,900] #TESTAR O NUMEROP DE PONTOS Devem ser muitos
ball_calibration= np.array(label_ball_calibration)+24
code_omni = [["MAX_H_ball_omni",0],
        ["MIN_H_ball_omni",0],
        ["MAX_S_ball_omni",0],
        ["MIN_S_ball_omni",0],
        ["MAX_V_ball_omni",0],
        ["MIN_V_ball_omni",0],
        ["MAX_H_field_omni",0],
        ["MIN_H_field_omni",0],
        ["MAX_S_field_omni",0],
        ["MIN_S_field_omni",0],
        ["MAX_V_field_omni",0],
        ["MIN_V_field_omni",0],
        ["MAX_H_lines_omni",0],
        ["MIN_H_lines_omni",0],
        ["MAX_S_lines_omni",0],
        ["MIN_S_lines_omni",0],
        ["MAX_V_lines_omni",0],
        ["MIN_V_lines_omni",0],
        ["MAX_H_robot_omni",0],
        ["MIN_H_robot_omni",0],
        ["MAX_S_robot_omni",0],
        ["MIN_S_robot_omni",0],
        ["MAX_V_robot_omni",0],
        ["MIN_V_robot_omni",0],
        ["MAX_H_redshirt_omni",0],
        ["MIN_H_redshirt_omni",0],
        ["MAX_S_redshirt_omni",0],
        ["MIN_S_redshirt_omni",0],
        ["MAX_V_redshirt_omni",0],
        ["MIN_V_redshirt_omni",0],
        ["MAX_H_blueshirt_omni",0],
        ["MIN_H_blueshirt_omni",0],
        ["MAX_S_blueshirt_omni",0],
        ["MIN_S_blueshirt_omni",0],
        ["MAX_V_blueshirt_omni",0],
        ["MIN_V_blueshirt_omni",0],
        ["EXPOSURE_omni",0],
        ["GAIN_omni",0],
        ["SATURATION_omni",0],
        ["MAX_AREA_ball_omni",0], #38
        ["MIN_AREA_ball_omni",0], #39
        
        ]
code_front = [["MAX_H_ball_front",0],
        ["MIN_H_ball_front",0],
        ["MAX_S_ball_front",0],
        ["MIN_S_ball_front",0],
        ["MAX_V_ball_front",0],
        ["MIN_V_ball_front",0],
        ["MAX_H_field_front",0],
        ["MIN_H_field_front",0],
        ["MAX_S_field_front",0],
        ["MIN_S_field_front",0],
        ["MAX_V_field_front",0],
        ["MIN_V_field_front",0],
        ["MAX_H_lines_front",0],
        ["MIN_H_lines_front",0],
        ["MAX_S_lines_front",0],
        ["MIN_S_lines_front",0],
        ["MAX_V_lines_front",0],
        ["MIN_V_lines_front",0],
        ["MAX_H_robot_front",0],
        ["MIN_H_robot_front",0],
        ["MAX_S_robot_front",0],
        ["MIN_S_robot_front",0],
        ["MAX_V_robot_front",0],
        ["MIN_V_robot_front",0],
        ["MAX_H_redshirt_front",0],
        ["MIN_H_redshirt_front",0],
        ["MAX_S_redshirt_front",0],
        ["MIN_S_redshirt_front",0],
        ["MAX_V_redshirt_front",0],
        ["MIN_V_redshirt_front",0],
        ["MAX_H_blueshirt_front",0],
        ["MIN_H_blueshirt_front",0],
        ["MAX_S_blueshirt_front",0],
        ["MIN_S_blueshirt_front",0],
        ["MAX_V_blueshirt_front",0],
        ["MIN_V_blueshirt_front",0],
        ]
code_warp = [["WARP_bl",0],["WARP_br",0]]



# Images of Buttons
image_pause = './aux_file/pause.png'
image_restart = './aux_file/refresh.png'
image_save = './aux_file/download.png'
image_play = './aux_file/play.png'
image_rec = './aux_file/rec-button1.png'
image_recording = './aux_file/rec-button.png'
selected_object=0
Calibration=0
Source=1 
nFrames=0
zeros = np.zeros((480,480,3), dtype=np.uint8)



real_dist_omni=np.zeros((480,480,2)) 
real_polar_omni=np.zeros((480,480,2)) 
a, b, c, d, e, f,g =0,0,0,0,0,0,0


real_dist_omni_ball=np.zeros((480,480,2)) 
real_polar_omni_ball=np.zeros((480,480,2)) 
a_ball, b_ball, c_ball, d_ball, e_ball, f_ball,g_ball =0,0,0,0,0,0,0

menu_def = [
            ['File', [ 'Save', 'Exit']],
            ['Camera to calibrate', ['Calibrate Front', 'Calibrate Omni']],
            ['Source', ['Camera', ['Front', 'Omni'], 'File',['Image','Video']]],
            ['Help', 'About...']]
var_to_config=[
    [sg.Button("->Ball", size=(10, 1),button_color=('black', 'yellow'), key="-Ball-"),],
    [sg.Button("Field", size=(10, 1),button_color=('black', 'green'), key="-Field-"),],[sg.Button("Lines", size=(10, 1),button_color=('black',"white"), key="-Lines-"),],
    [sg.Button("Robot", size=(10, 1),button_color=('white', 'black'), key="-Robot-"),],[sg.Button("Red Shirt", size=(10, 1),button_color=('white', 'red'), key="-Rshirt-"),],
    [sg.Button("Blue Shirt", size=(10, 1),button_color=('white', 'blue'), key="-Bshirt-"),]]
hsv_config=[
        [sg.Text("Max. H",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black" ,
            #color="black",
            key="-MAXH-",
        ),],
        [sg.Text("Min. H",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+1][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MINH-",
        ),],
        [sg.Text("Max. S",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+2][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MAXS-",
        ),],
        [sg.Text("Min. S",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+3][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MINS-",
        ),],
        [sg.Text("Max. V",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+4][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MAXV-",
        ),],
        [sg.Text("Min. V",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+5][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MINV-",
        ),],
        [sg.Text("Expo",background_color="white",text_color="black"),
        sg.Slider(
            (456, 327540),#Min 45.6 Max 32754.0
            460,
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-EXP-",
        ),],
        [sg.Text("Gain",background_color="white",text_color="black"),
        sg.Slider(
            (0, 239), # Min 0 Max 23.9 
            200,
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-GAIN-",
        ),],
        [sg.Text("Satur",background_color="white",text_color="black"),
        sg.Slider(
            (0, 3999), #Min 0 Max 399.9
            200,
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-SAT-",
        ),],
        ]
hsv_config2=[
        [sg.Text("Max. H",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MAXH2-",
        ),],
        [sg.Text("Min. H",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+1][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MINH2-",
        ),],
        [sg.Text("Max. S",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+2][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MAXS2-",
        ),],
        [sg.Text("Min. S",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+3][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MINS2-",
        ),],
        [sg.Text("Max. V",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+4][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MAXV2-",
        ),],
        [sg.Text("Min. V",background_color="white",text_color="black"),
        sg.Slider(
            (0, 255),
            code_omni[selected_object*6+5][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-MINV2-",
        ),],
        [sg.Text("Offset X",background_color="white",text_color="black"),
        sg.Slider(
            (0, 480),
            code_omni[selected_object*6+5][1],
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-OFF2-",
        ),],]
warp_config=[
    [sg.Text("Warp",background_color="white",text_color="black"),
        sg.Slider(
            (0, 240),
            code_omni[selected_object*6+5][1],
            1,
            orientation="h",
            size=(110, 15),
            background_color="white",
            text_color="black",
            key="-WARP-",
        ),],]
field_config=[
    [sg.Text("A = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_A-"),sg.Text("B = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_B-"),sg.Text("C = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_C-")],
    [sg.Text("D = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_D-"),sg.Text("E = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_E-"),sg.Text("F = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_F-")],
    [sg.Text("G = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_G-"),sg.Text("H = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_H-"),sg.Text("I = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_I-")],
    [sg.Text("J = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_J-"),sg.Text("K = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_K-"),sg.Text("L = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_L-")],
    [sg.Text("M = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_M-"),sg.Text("N = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_N-"),sg.Text("O = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_O-")],
    [sg.Text("P = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_P-"),sg.Text("Q = ", size=(4, 1),text_color="black",background_color="white"),sg.InputText(size=(15,1),key="-F_Q-")],
    
    ]
#layout1 = [[, sg.Column(col2, element_justification='c')
#      , sg.Column(col3, element_justification='c'), sg.Column(col4, element_justification='c')]]
# Define the window layout
layout1 = [
    [sg.Text("Original", size=(68, 1), justification="center"),sg.Text("Filtred", size=(68, 1), justification="center")],
    [sg.Image(filename="", key="-IMAGE_ORI-"),sg.Image(filename="", key="-IMAGE-")],
    [sg.Column(var_to_config, element_justification='c' ),sg.Column(hsv_config, element_justification='c',background_color='white')],
    [sg.Button("Save Constants", size=(30, 1),button_color=('black', 'white'),key="-SAVE_CONST-")]]
layout2 = [
    [sg.Text("Original", size=(68, 1), justification="center"),sg.Text("Map", size=(68, 1), justification="center")],
    [sg.Image(filename="", key="-IMAGE_ORI2-"),sg.Image(filename="", key="-IMAGE2-")],
    [sg.Column(hsv_config2, element_justification='c',background_color='white' ),sg.Image(filename="", key="-IMAGE_PLOT-")],
    [ sg.Text("Equation: y = ax⁶ + bx⁵ + cx⁴ + dx³ + ex² + fx + g",size=(200,1),text_color="white",key="-Equation-"),],
    [sg.Button("Save Distances", size=(30, 1),button_color=('black', 'white'),key="-SAVE_DIST-")]]   
layout3 = [
    [sg.Text("Original", size=(68, 1), justification="center"),sg.Text("Warp", size=(68, 1), justification="center")],
    [sg.Image(filename="", key="-IMAGE_ORI3-"),sg.Image(filename="", key="-IMAGE3-")],
    [sg.Column(warp_config, element_justification='c',background_color='white' )],
    ]
layout4 = [
    [sg.Text("Field", size=(68, 1), justification="center")],
    [sg.Image(filename="./aux_file/Field_Original.png")],
    [sg.Column(field_config, element_justification='c',background_color='white' )],
    [sg.Button("Save Field Distances", size=(30, 1),button_color=('black', 'white'),key="-SAVE_FIELD-")],
    ]
layout5 = [
    [sg.Text("Original", size=(68, 1), justification="center")],
    [sg.Image(filename="", key="-IMAGE_CENTER-")],
    ]
layout6 = [
    [sg.Text("Original", size=(68, 1), justification="center"),sg.Text("Map", size=(68, 1), justification="center")],
    [sg.Image(filename="", key="-IMAGE_ORI4-"),sg.Image(filename="", key="-IMAGE4-"),sg.Image(filename="", key="-IMAGE_PLOT2-")],
     [sg.Text("Max. Area",background_color="white",text_color="black"),
        sg.Slider(
            (0, 4000),
            2000,
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-Max_Area-",
        ),],
        [sg.Text("Min. Area",background_color="white",text_color="black"),
        sg.Slider(
            (0, 4000),
            0,
            1,
            orientation="h",
            size=(90, 15),
            background_color="white",
            text_color="black",
            key="-Min_Area-",
        ),],
    [ sg.Text("Equation: y = ax⁶ + bx⁵ + cx⁴ + dx³ + ex² + fx + g",size=(200,1),text_color="white",key="-Equation_Ball-"),],
    [ sg.Text("Put Ball at "+ str(label_ball_calibration[0]/100)+" Meters to the center of the robot",size=(200,5),text_color="white",key="-Put_ball-"),],
    [sg.Button("Check", size=(30, 1),button_color=('black', 'white'),key="-SAVE_BALL-")],
    [sg.Button("Save Distances", size=(30, 1),button_color=('black', 'white'),key="-SAVE_DIST_BALL-")],
    [ sg.Text("X= 0 Y= 0 D= 0 A= 0",size=(200,1),text_color="white",key="-Coordenates_Ball-")],] 
      
tabgrp = [
    [sg.Menu(menu_def,text_color='black', font="SYSTEM_DEFAULT", pad=(10,10))],
        [sg.Text("Calibrate: ",text_color="black",background_color='firebrick4',font=('Any 15')),sg.Text("Omni",text_color="black",background_color='firebrick4',font=('Any 15'),key="-CALIBRATE-"),
        sg.Text("     Source: ",text_color="black",background_color='firebrick4',font=('Any 15')),sg.Text("Image",text_color="black",background_color='firebrick4',font=('Any 15'),key="-SOURCE-"),
        sg.Text("                    ",background_color='firebrick4'),
        sg.Button("    ", image_size=(30, 30), button_color=("firebrick4","firebrick4"),image_filename=image_pause, image_subsample=4, border_width=0,key="Pause"),
        sg.Button("    ", image_size=(40, 40), button_color=("firebrick4","firebrick4"),image_filename=image_play, image_subsample=4, border_width=0,key="Continue"),
        sg.Button("    ", image_size=(30, 30), button_color=("firebrick4","firebrick4"),image_filename=image_restart, image_subsample=4, border_width=0,key="Restart"),
        sg.Button("    ", image_size=(30, 30), button_color=("firebrick4","firebrick4"),image_filename=image_rec, image_subsample=4, border_width=0,key="-Record-"),
        sg.Button("    ", image_size=(30, 30), button_color=("firebrick4","firebrick4"),image_filename=image_save, image_subsample=4, border_width=0,key="Save Record"),
        sg.Text("  Name of video file to save: ",text_color="black",background_color='firebrick4'),
        sg.InputText("My_new.avi",size=(15,1),key="-SAVE_V-"),
        sg.Image(filename="./aux_file/logo_LAR.png",background_color="firebrick4", size=(300,50))],
        [sg.TabGroup([[
            sg.Tab('Center Calibration', layout5, title_color=('Black'), border_width =100,tooltip='Center Calibration', element_justification= 'center',key="-TAB_CENTER-"),
            sg.Tab('HSV Calibration', layout1, title_color=('Black'), border_width =100,tooltip='HSV Calibration', element_justification= 'center',key="-TAB_HSV-"),
            sg.Tab('Map Calibration', layout2,title_color='Blue', element_justification= 'center',key="-TAB_MAP-"),
            sg.Tab('Ball Calibration', layout6,title_color='Blue', element_justification= 'center',key="-TAB_BALL-"),
            sg.Tab('Warp', layout3,title_color='Black',tooltip='Warp', element_justification= 'center',key="-TAB_WARP-"),
            sg.Tab('Field Calibration', layout4, title_color=('Black'), border_width =100,tooltip='Field Calibration', element_justification= 'center',key="-TAB_FIELD-")
            
            ]],key='_TAB_GROUP_', enable_events=True,tab_location='centertop',title_color='white',background_color='firebrick4', tab_background_color='black',selected_title_color='black',selected_background_color='white', border_width=5)],
        ]  

