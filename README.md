# go_pca9685
I bought an Adafruit Servo Hat - it works totally differently to the Pimoroni Pantilt hat!

This is my Go port of the Adafruit pca9685 library for the Adafruit panTilt hat on a RasPi

Didn't get on with the experimental pca9685 device package under periph... so looked at other libraries.
please refer to this:-

https://learn.adafruit.com/adafruit-16-channel-pwm-servo-hat-for-raspberry-pi/library-reference

Note that the values of servo Max and Min are different to those in periph.
Also the experimental code in periph needs a stop/halt function!

NBBB I haven't gone into the depths of PWM - The first attempts managed to turn my little servos into 360degree rotators
(there was a sound of the end stop being destroyed in at least one of them)
Amazingly they still work now that the software is under control.

Now where were those Ping-Pong balls to make the eyes ....
RichR
