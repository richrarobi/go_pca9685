// pca9685
// Rich Robinson
// Dec 2018
// This is a Go port of the Adafruit Python library
//
package pca9685

import (
//    "fmt"
    "log"
    "time"
//    "reflect"
//    "strconv"
    "periph.io/x/periph/host"
//    "periph.io/x/periph/host/rpi"
    "periph.io/x/periph/conn/i2c/i2creg"
    "periph.io/x/periph/conn/i2c"
)

var dev i2c.Dev
var bus i2c.BusCloser
const pca9685Addr uint16 = 0x40
const i2cbus string = "1"
const confreg byte = 0x00

const degMin int = -90
const degMax int =  90

const swrst    byte = 0x06

// pca9865 registers
const mode1    byte = 0x00
const mode2    byte = 0x01
const subaddr1 byte = 0x02
const subaddr2 byte = 0x03
const subaddr3 byte = 0x04
const prescale byte = 0xfe
const outxonl  byte = 0x06
const outxonh  byte = 0x07
const outxoffl byte = 0x08
const outxoffh byte = 0x09
const noutonl  byte = 0xfa
const noutonh  byte = 0xfb
const noutoffl byte = 0xfc
const noutoffh byte = 0xfd

// pca9865 bits
const allcall  byte = 0x01
const invrt    byte = 0x10
const sleep    byte = 0x10
const restart  byte = 0x80
const outdrv   byte = 0x04

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func Open() {
    var err error
// initialise periph
        if _, err := host.Init(); err != nil {
            log.Fatal(err)
    }
        bus, err = i2creg.Open( i2cbus )
        if err != nil {
            log.Fatal(err)
    }
//    fmt.Println("opened ok: ", bus, reflect.TypeOf(bus))
    dev = i2c.Dev{bus, pca9685Addr}
//    fmt.Println("dev: ", dev, reflect.TypeOf(dev))
    SetAllPWM(0,0)
    i2cWriteByte(mode2, outdrv)
    i2cWriteByte(mode1, allcall)
    delay(50)
    m1 := i2cReadByte(mode1)
    m1 = m1 & ^sleep
    i2cWriteByte(mode1, m1)
    delay(50)
}

func Close() {
    i2cWriteByte(noutoffh, 0x10)
    delay(50)
    bus.Close()
}

func SetPWMFreq(hz int) {
    var pl float32 = 25000000 / 4096
    pl /= float32(hz)
    pl -= 1.0
//    fmt.Println("setPWMFreq: hz: ", hz)
//    fmt.Println("setPWMFreq: pl: ", pl)
    pscale := byte(pl + 0.5)
//    fmt.Println("setPWMFreq: pscale: ", pscale)

    old := i2cReadByte(mode1)
    i2cWriteByte(mode1, (old & 0x7f) | 0x10) // sleep
    i2cWriteByte(prescale, pscale)
    i2cWriteByte(mode1, old)
    delay(50)
    i2cWriteByte(mode1, old | 0x80)
}

func SetPWM(out int, on uint16, off uint16) {
    i2cWriteByte( outxonl + byte(4 * out), byte(on))
    i2cWriteByte( outxonh + byte(4 * out), byte(on >> 8))
    i2cWriteByte( outxoffl + byte(4 * out), byte(off))
    i2cWriteByte( outxoffh + byte(4 * out), byte(off >> 8))
}

func SetAllPWM(on uint16, off uint16) {
    i2cWriteByte( noutonl, byte(on))
    i2cWriteByte( noutonh, byte(on >> 8))
    i2cWriteByte( noutoffl, byte(off))
    i2cWriteByte( noutoffh, byte(off >> 8))
}

func i2cReadByte(reg byte) (res byte) {
    read := make([]byte, 1)
    write := []byte{reg}
    if err := dev.Tx(write, read); err != nil {
        log.Fatal(err)
    }
    res = read[0]
    return res
}

func i2cWriteByte(reg byte, data byte) {
    read := make([]byte, 0)
    write := []byte{reg}
    write = append(write, data)
// write to i2c
    if err := dev.Tx(write, read); err != nil {
        log.Fatal(err)
    }
}
