#!/bin/bash

ffmpeg -r 60 image2 -s 1920x1080 -i color%06d.png -vcodec libx264 -crf 25 -pix_fmt yuv420p color.mp4