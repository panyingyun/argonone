package dev

import (
	"argonone/log"
	"io/ioutil"
	"strconv"
	"strings"
)

//https://github.com/Elrondo46/argonone/branches
//CPU temperature: cat /sys/class/thermal/thermal_zone0/temp
//example: 46251

const (
	cputemp = "/sys/class/thermal/thermal_zone0/temp"
)

type CPUTemp struct {
	name         string
	temperate    float64
	temperateint int64
}

func NewCPUTemp() *CPUTemp {
	return &CPUTemp{
		name:         "argonone",
		temperate:    0.0,
		temperateint: 0,
	}
}

func (d *CPUTemp) FetchTemperate() (err error) {
	data, err := ioutil.ReadFile(cputemp)
	if err != nil {
		log.Default().Infof("fetch temperature fail. err: %v", err)
		return
	}
	ret := strings.ReplaceAll(string(data), "\n", "")
	temp, err := strconv.ParseInt(string(ret), 10, 64)
	if err != nil {
		log.Default().Infof("ParseFloat  fail. err: %v", err)
		return
	}
	d.temperate = float64(temp) / 1000.0
	d.temperateint = temp
	return
}

func (d *CPUTemp) Name() string {
	return d.name
}

func (d *CPUTemp) Temperate() float64 {
	return d.temperate
}

func (d *CPUTemp) TemperateInt() int64 {
	return d.temperateint
}
