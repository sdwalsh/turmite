# Turmites #

Turmites is a small Go program that runs one more more [turmites](https://en.wikipedia.org/wiki/Turmite) through a mound.

Mound library located [here](mound)

Every turmite is a 2d turing machine, one of the most famous turmites being [Langton's Ant](https://en.wikipedia.org/wiki/Langton%27s_ant) (a four state turing machine).

![turmite moving around a mound](https://raw.githubusercontent.com/sdwalsh/turmites/master/color.gif)

## How do I get set up? ##

* Linux/Unix-like OS required
* `bash`, `rm`, `ffmpeg` required for `func mound.BatchImages(...)`

## Things to consider ##

* Currently every frame is encoded into png files into a temporary directory. This takes a long time!
* `mound.BatchImages(...)` currently handles a max of 999,999 frames
* Want to convert .mp4 to .gif? Consider [this](https://superuser.com/questions/556029/how-do-i-convert-a-video-to-gif-using-ffmpeg-with-reasonable-quality) Stack Overflow post
* API is not stable and subject to change

## TODO ##

* Modify `mound.BatchImages(...)` to start generating pngs at a specific step (currently steps from frame 1 -> frame 999,999)
* Create function to merge various .mp4s
* Allow skipping of frames
* Allow user to choose output type (currently only allows .mp4)
* Consider pushing png encoding to different cores (currently the bottleneck)

## Who do I talk to? ##

* Contact Sean @ [github.com/sdwalsh](https://www.github.com/sdwalsh) or [bitbucket.org/sdwalsh](https://www.bitbucket.org/sdwalsh)

## Contributing ##

Feel free to fork this repository and send pull requests!

# License #

Copyright 2017 Sean Walsh

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.