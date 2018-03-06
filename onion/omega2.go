package onion

import (
	"fmt"
	"log"

	"github.com/DazWilkin/WebThingGolang/webthing"
)

type Device struct {
	Name     string
	MAC      string
	Password string
	CloudID  string
}

func NewDevice(name, mac, password, cloudid string) (Device, error) {
	d := Device{
		Name:     name,
		MAC:      mac,
		Password: password,
		CloudID:  cloudid}
	return d, nil
}
func (d *Device) String() string {
	return fmt.Sprintf("%v\t%v\t%v", d.MAC, d.Password, d.CloudID)
}
func (d *Device) ToWebThing() webthing.WebThing {
	log.Println("[Device:toWebThing]")
	wt := webthing.WebThing{
		Name: d.Name,
		Type: "Omega2"}
	return wt
}
