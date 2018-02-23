package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"

	particle "github.com/DazWilkin/webthing/particle"
	webthing "github.com/DazWilkin/webthing/webthing"
)

var (
	_      particle.Device
	_      webthing.WebThing
	random *rand.Rand
)

type NameIDPair struct {
	Name string
	ID   string
}

// Generate random seed
func init() {
	source := rand.NewSource(time.Now().UnixNano())
	random = rand.New(source)
}

// Avoid including Particle-specific data in code
// Expects DEVICES=device1=id1,...
func getDevicesFromEnv() []NameIDPair {
	result := []NameIDPair{}
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to find '.env' proceeding")
	}
	envDevices := os.Getenv("DEVICES")
	if len(envDevices) == 0 {
		log.Fatal("Expect 'DEVICES=device1=id1,device2=id2,...'")
	}
	devicesIDName := strings.Split(envDevices, ",")
	for _, deviceIDName := range devicesIDName {
		IDName := strings.Split(deviceIDName, "=")
		result = append(result, NameIDPair{
			Name: IDName[0],
			ID:   IDName[1]})
	}
	return result
}

// Generate a stringified (!) random float 30±10
func randomStr() string {
	return strconv.FormatFloat(random.Float64()*10+30, 'f', 8, 64)
}

func main() {
	log.Println("[Main]")

	e := getDevicesFromEnv()

	devices := []particle.Device{}

	// Three way to create Devices
	// 1. By inlining the structs
	// 2. Using "New" funcs for parent and child strucs
	// 3. Using "New" func for parent but not children

	// Create Device
	// Uses
	// This doesn't include the since-added Functions
	devices = append(devices, particle.Device{
		ID:        e[0].ID,
		Name:      e[0].Name,
		Type:      "Photon",
		Firmware:  "0.7.0-rc.1",
		Handshake: time.Now(),
		Variables: map[string]particle.Variable{
			"temperature": particle.Variable{
				Type:  "float",
				Value: randomStr()}}})

	var device particle.Device

	device, _ = particle.NewDevice(
		e[1].ID,
		e[1].Name,
		"Photon",
		"0.7.0-rc.1",
		time.Now(),
		make(map[string]particle.Function),
		make(map[string]particle.Variable))
	device.AddVariable("temperature", particle.Variable{Type: "float", Value: randomStr()})
	device.AddVariable("temperature", particle.Variable{Type: "float", Value: randomStr()})
	device.AddVariable("humidity", particle.Variable{Type: "float", Value: randomStr()})
	devices = append(devices, device)

	device, _ = particle.NewDevice(
		e[2].ID,
		e[2].Name,
		"Raspberry Pi",
		"unknown",
		time.Now(),
		map[string]particle.Function{},
		map[string]particle.Variable{},
	)
	devices = append(devices, device)

	// Add the Devices to an aggregating Account
	// Account provides no functional value beyond this glue
	account, _ := particle.NewAccount("someone@gmail.com", []particle.Device{})
	for _, device := range devices {
		account.AddDevice(device)
	}
	fmt.Printf("a: %v\n", account.String())

	// Enumerate the devices
	// Convert to String
	// Convert to a preliminary (!) WebThing
	for _, device := range account.Devices {
		fmt.Printf("device: %v\n", device.String())
		fmt.Printf("device: %v\n", device.ToWebThing())
	}

}
