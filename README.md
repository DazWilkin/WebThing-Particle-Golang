# Web Thing API For Particle Devices in Golang

## Synopsis

To develop my [Golang](golang.org) skillz thought I'd try to combine an interest in IoT and a 'herd' of [Particle](https://www.particle.io/) devices into an attempt at a proof-of-concept (!) SDK for Mozilla's [Web Thing API](https://mozilla-iot.github.io/wot/)

## Build

* Clone the library using `git clone` or `go get`
* Either set an environment variable `DEVICES=device1=id1,device2=id2...` or create a `.env` file containing this value
* run it

```bash
git clone https://github.com/DazWilkin/WebThingGolang.git
cd WebThingGolang
# Replace deviceX and idX with your Particle device names, ids if you have them
echo "DEVICES=device1=id1,device2=id2" > .env
go run main.go
```

or in your working directory:

```bash
mkdir go
GOPATH=$PWD/go
PATH=$PATH:$GOPATH/bin
go get github.com/DazWilkin/WebThingGolang
cd go/src/github.com/DazWilkin/WebThingGolang
# Replace deviceX and idX with your Particle device names, ids if you have them
echo "DEVICES=device1=id1,device2=id2" > .env
go run main.go
```

## Status

I have a trivial temperature and humidity sensor configured on several of my Particles. The `main.go` configures my devices with the `temperature` and `humidity` sensor variables. Configure your particles as you wish.

Liberal logging:
```
2018/02/23 12:39:06 [Main]
...
a: someone@gmail.com {123456789012345678901234, 123456789012345678901234, }
2018/02/23 12:39:06 [Device:String]
2018/02/23 12:39:06 [Device:String:λVariables] Entered
2018/02/23 12:39:06 temperature={float 39.56433425}
2018/02/23 12:39:06 [Device:String:λFunctions] Entered
device: [123456789012345678901234] name0 functions: {} variables: {temperature: {float 39.56433425},}
2018/02/23 12:39:06 [Device:toWebThing]
device: {tasty_doctor Photon  map[temperature:{float   /v1/devices/123456789012345678901234/temperature}]}
2018/02/23 12:39:06 [Device:String]
2018/02/23 12:39:06 [Device:String:λVariables] Entered
2018/02/23 12:39:06 temperature={float 36.71775245}
2018/02/23 12:39:06 humidity={float 36.74877121}
2018/02/23 12:39:06 [Device:String:λFunctions] Entered
device: [123456789012345678901234] name1 functions: {} variables: {temperature: {float 36.71775245},humidity: {float 36.74877121},}
2018/02/23 12:39:06 [Device:toWebThing]
device: {aardvark_scrapple Photon  map[humidity:{float   /v1/devices/123456789012345678901234/humidity} temperature:{float   /v1/devices/123456789012345678901234/temperature}]}
...
2018/02/23 12:39:06 [serveThings:go] name3→:40003
2018/02/23 12:39:06 [serveThings:go] name1→:40001
2018/02/23 12:39:06 [serveThings:go] name2→:40002
2018/02/23 12:39:06 [serveThings:go] name4→:40004
2018/02/23 12:39:06 [serveThings:go] name0→:40000
```

Ultimately, each device is exposed on its own port and provides **basic** Web Thing output:
```JSON
curl http://localhost:40000
{
    "name":"name0",
    "type":"Photon",
    "properties":{
        "temperature":{
            "type":"float",
            "href":"/v1/devices/2b0039000247343337373738/temperature"
        }
    }
}
```