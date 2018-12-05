// Rich Robinson
// Dec 2018
//
package main

import (
    "os"
    "os/signal"
    "syscall"
    "fmt"
    "log"
    "time"
    sv "github.com/richrarobi/pca9685"
)

// note these values !!
const servoMin uint16 = 150
const servoMax uint16 = 600

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
// initialise getout
    running := true
    signalChannel := make(chan os.Signal, 2)
    signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
    go func() {
        sig := <-signalChannel
        switch sig {
        case os.Interrupt:
            log.Println("Stopping on Interrupt")
            running = false
            return
        case syscall.SIGTERM:
            log.Println("Stopping on Terminate")
            running = false
            return
        }
    }()
    sv.Open()
    sv.SetPWMFreq(60)
    for running {
        sv.SetPWM(0,0, servoMin)
        sv.SetPWM(3,0, 0)           // this is an led
        sv.SetPWM(4,0, servoMin)
        sv.SetPWM(7,0, servoMin)
        sv.SetPWM(11,0, servoMin)
        sv.SetPWM(15,0, servoMin)
        delay(1000)
        sv.SetPWM(0,0, servoMax)
        sv.SetPWM(3,0, servoMax)
        sv.SetPWM(4,0, servoMax)
        sv.SetPWM(7,0, servoMax)
        sv.SetPWM(11,0, servoMax)
        sv.SetPWM(15,0, servoMax)
        delay(1000)
    }
        delay(1000)
        sv.SetPWM(0,0, servoMin)
        sv.SetPWM(3,0, 0)
        sv.SetPWM(4,0, servoMin)
        sv.SetPWM(7,0, servoMin)
        sv.SetPWM(11,0, servoMin)
        sv.SetPWM(15,0, servoMin)
        delay(1000)
    sv.Close()
    fmt.Println("Leaving Program")
}
