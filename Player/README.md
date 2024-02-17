![LAR](https://github.com/MSL-LAR-MinhoTeam/2TDP/blob/main/Images/git_msl_player.png)
# LAR@MSL Player Code

## Why Golang?
Main reasons:
- Fast;
- Easy to learn;
- Simple thread system.

## Communication
Communication is used like a local database. Communicates with ESP32 (low level) and BaseStation (high level).

## OmniVision_pkg
Use only the omnidirectional camera to self-localization and objects detection.

## Freenect
Library to adquire images from Kinect.

## Skills
Use kinect camera as primary data source and performe all individual skills.

## GoalKeeper Vision
GoalKeeper Vision will be responsible for the prediction of the ball trajectory and the response of the GoalKeeper to defend the goal.

# Run
## Run Calibration
> $ make calib

## Run Server 
Just once, if server dont interrupted.
> $ make server

## Run Go
> $ make debug

## Run All
> $ make start
