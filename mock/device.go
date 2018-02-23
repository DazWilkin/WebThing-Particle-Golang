package mock

import (
	"fmt"
	"log"

	"github.com/DazWilkin/WebThing-Particle-Golang/webthing"
)

type Device struct {
	Name string
	Type string
}

func NewDevice(name, _type string) Device {
	log.Println("[NewDevice] Entered")
	d := Device{
		Name: name,
		Type: _type}
	return d
}
func (d *Device) String() string {
	log.Println("[String] Entered")
	return fmt.Sprintf("%v [%v]\n", d.Name, d.Type)
}
func (d *Device) ToWebThing() webthing.WebThing {
	log.Println("[Device:toWebThing]")
	wt := webthing.WebThing{
		Name: d.Name,
		Type: d.Type}
	return wt
}
