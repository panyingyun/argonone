# PI4-CASE-ARGON-ONE
Argon One Service and Control Service  coding by golang  for Raspberry Pi 4B/4B+ centos7.7 or other Pi Linux System

Enjoy it.

### 1、Hardware 
- Raspberry Pi 4 Model B Rev 1.2
- PI4-CASE-ARGON-ONE

### 2、Rom & OS 
- CentOS 7.7 [Download Rom and Flash to SD](!https://www.michaelapp.com/posts/2018/2018-09-13-%E6%A0%91%E8%8E%93%E6%B4%BE64%E4%BD%8DCentOS%E5%AE%89%E8%A3%85/)

### 3、Boot config file(very important,if not do this, application can not run normal.)

*open i2c dev driver*

```shell
### vim /boot/config.txt and add next info
dtparam=i2c_arm=on
dtparam=i2c-1=on

### vim /etc/modules-load.d/modules.conf and add next info
i2c-dev
i2c-bcm2835
i2c-bcm2708

### vim /etc/modprobe.d/custom.conf and add next info
### Very Important Config Here(Change I2C Baudrate)
### If we use default baudrate(110K)，i2cdetect can not find ARGON-ONE
options i2c_bcm2708 baudrate=32000
options i2c_bcm2835 baudrate=32000
```

### 4、Reboot && check config work or not 

 - (1) check i2c dirver work 

```shell
[root@pinas2 modprobe.d]# lsmod | grep i2c
i2c_bcm2708            16384  0 
i2c_bcm2835            16384  0 
i2c_dev                20480  0 
```
 - (2) check i2c baudrate(32000 is OK)

```shell
[root@pinas2 tools]# cat /sys/module/i2c_bcm2708/parameters/baudrate
32000
```

 - (3) check ARGON-ONE device('1a' is device address for fan) 

```shell
[root@pinas2 tools]# yum install i2c-tools -y
[root@pinas2 tools]# i2cdetect -y 1 
     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
00:          -- -- -- -- -- -- -- -- -- -- -- -- -- 
10: -- -- -- -- -- -- -- -- -- -- 1a -- -- -- -- -- 
20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
40: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
70: -- -- -- -- -- -- -- --  
```

- (4) check cpu temp 

```shell
[root@pinas2 tools]# cat /sys/class/thermal/thermal_zone0/temp
44790
```

### 5、How to build 

```shell
//ssh to your board and prepare Pioneer600
//and install git tools 
sudo apt-get update
sudo apt-get upgrade
sudo apt install -y vim git build-essential i2c-tools
sudo apt install -y golang

//clone code 
git clone git@github.com:panyingyun/argonone.git

//build 
cd argonone/src/cmd

//build arm
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build  -o argonone
```


### 6、How to run 

```shell
argonone -c prod.yml
```

### 7、Auto-run when reboot OS 

Run build_arm64.sh, it will autorun argonone when reboot os

```shell
chmod 755 build_arm64.sh 
sudo sh build_arm64.sh 
```

Other shell cmd 

```shell
// enable 
sudo systemctl enable argonone.service

// disable 
sudo systemctl disable argonone.service

// start service 
sudo systemctl start  argonone.service

// stop 
sudo systemctl stop argonone.service

// restart 
sudo systemctl restart argonone.service
```

### 8、Check Result

**When CPU Temp from 41.381℃ to 39.92℃ ,and fan speed from 10% to stop.**

```shell
{"level":"INFO","ts":"2020-06-07 11:20:34","msg":"Raspberry Pi 4 Argonone Fan"}
{"level":"INFO","ts":"2020-06-07 11:20:34","msg":"Thanks to https://gobot.io"}
{"level":"INFO","ts":"2020-06-07 11:20:35","msg":"Current Temp is 41.381 ℃"}
{"level":"INFO","ts":"2020-06-07 11:20:35","msg":"FANOn10 ..."}
{"level":"INFO","ts":"2020-06-07 11:20:40","msg":"Current Temp is 40.894 ℃"}
{"level":"INFO","ts":"2020-06-07 11:20:40","msg":"FANOn10 ..."}
{"level":"INFO","ts":"2020-06-07 11:20:45","msg":"Current Temp is 39.92 ℃"}
{"level":"INFO","ts":"2020-06-07 11:20:45","msg":"FANOff ..."}
{"level":"INFO","ts":"2020-06-07 11:20:50","msg":"Current Temp is 39.92 ℃"}
{"level":"INFO","ts":"2020-06-07 11:20:50","msg":"FANOff ..."}
```

### 9、Thanks To

- https://github.com/Elrondo46/argonone
- http://www.waveshare.net/wiki/PI4-CASE-ARGON-ONE
- https://www.cnblogs.com/bcy520/p/6816310.html

