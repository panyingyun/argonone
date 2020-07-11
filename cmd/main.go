package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"argonone/dev"
	"argonone/log"

	cron "github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

const (
	//Spec:https://www.cnblogs.com/wangyuyu/p/4230742.html
	cronSpec     = "@hourly"     //Every hour
	cronSpecTest = "*/2 * * * *" //Every minute
)
var fan *dev.Fan
var cpu *dev.CPUTemp
//CPU温度	风扇转速
//45度	50%
//60度	100%
func CheckCPUTempAndFanStatus() {
	err := cpu.FetchTemperate()
	if err != nil {
		log.Default().Error("err = ", err)
		return
	}
	temp := cpu.TemperateInt()
	log.Default().Infof("Current Temp is %v ℃", cpu.Temperate())
	switch {
	case temp > 60000:
		_ = fan.FANOn100()
	case temp > 50000:
		_ = fan.FANOn50()
	default:
		_ = fan.FANOff()
	}
}

func run(c *cli.Context) error {

	fmt.Println("conf = ", c.String("conf"))
	config := viper.New()
	config.SetConfigFile(c.String("conf"))
	config.SetConfigType("yaml")
	config.ReadInConfig()
	opt, err := log.NewOptions(config)
	if err != nil {
		fmt.Println("err = ", err)
	}
	logger, err := log.NewLogger(opt)
	if err != nil {
		fmt.Println("err = ", err)
	}
	logger.Info("Raspberry Pi 4 Argonone Fan")
	logger.Info("Thanks to https://gobot.io")
	fan = dev.NewFan()
	cpu = dev.NewCPUTemp()
	cron := cron.New()
	cron.AddFunc(cronSpecTest, CheckCPUTempAndFanStatus)
	cron.Start()
	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	logger.Infof("signal received signal %v", <-sigChan)
	logger.Warn("shutting down server")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "argonone"
	app.Usage = "/usr/bin/argonone -c /etc/argonone/prod.yml"
	app.Version = "0.0.1"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "conf,c",
			Usage:  "Set conf path here",
			Value:  "prod.yml",
			EnvVar: "APP_CONF",
		},
	}
	app.Run(os.Args)
}
