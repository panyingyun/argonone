package dev

import (
	"argonone/driver"
	"argonone/log"
)

//=======================
//CPU温度	风扇转速
//40度	10%
//50度	50%
//60度	100%
//=======================

const (
	I2cAddrFan  int    = 0x1a
	I2cFanOn100 byte   = 0x64
	I2cFanOn50  byte   = 0x32
	I2cFanOn10  byte   = 0x0a
	I2cFanOff   byte   = 0x00
	I2cFan      string = "/dev/i2c-1"
)

type Fan struct {
	i2c *driver.I2CDevice
}

func NewFan() *Fan {
	dev, err := driver.NewI2cDevice(I2cFan)
	if err != nil {
		log.Default().Errorf("NewI2cDevice err: ", err)
		return nil
	}
	err = dev.SetAddress(I2cAddrFan)
	if err != nil {
		log.Default().Errorf("SetAddress err: ", err)
		return nil
	}
	return &Fan{
		i2c: dev,
	}
}

func (p *Fan) FANOn100() error {
	log.Default().Info("FANOn100 ...")
	return p.i2c.WriteByteData(0x00, I2cFanOn100)
}

func (p *Fan) FANOn50() error {
	log.Default().Info("FANOn50 ...")
	return p.i2c.WriteByteData(0x00, I2cFanOn50)
}

func (p *Fan) FANOn10() error {
	log.Default().Info("FANOn10 ...")
	return p.i2c.WriteByteData(0x00, I2cFanOn10)
}

func (p *Fan) FANOff() error {
	log.Default().Info("FANOff ...")
	return p.i2c.WriteByteData(0x00, I2cFanOff)
}
