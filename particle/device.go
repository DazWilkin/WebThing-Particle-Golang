package particle

import (
	"fmt"
	"log"
	"time"

	"github.com/DazWilkin/WebThing-Particle-Golang/webthing"
)

type Device struct {
	ID        string
	Name      string
	Type      string
	Firmware  string
	Handshake time.Time
	Functions map[string]Function
	Variables map[string]Variable
}

func NewDevice(
	id, name, _type, firmware string,
	handshake time.Time,
	functions map[string]Function,
	variables map[string]Variable) (Device, error) {
	log.Println("[NewDevice] Entered")
	d := Device{
		ID:        id,
		Name:      name,
		Type:      _type,
		Firmware:  firmware,
		Handshake: handshake,
		Functions: functions,
		Variables: variables}
	return d, nil
}
func (d *Device) AddFunction(name string, function Function) {
	log.Println("[AddFunction] Entered")
	d.Functions[name] = function
}
func (d *Device) AddVariable(name string, variable Variable) {
	log.Println("[AddVariable] Entered")
	d.Variables[name] = variable
}
func (d *Device) String() string {
	log.Println("[Device:String]")
	variables := func(variables map[string]Variable) string {
		log.Println("[Device:String:λVariables] Entered")
		result := ""
		for key, value := range variables {
			log.Printf("%v=%v\n", key, value)
			result += fmt.Sprintf("%v: %v,", key, value)
		}
		return result
	}(d.Variables)
	functions := func(functions map[string]Function) string {
		log.Println("[Device:String:λFunctions] Entered")
		result := ""
		for key, value := range functions {
			log.Printf("%v=%v\n", key, value)
			result += fmt.Sprintf("%v: %v", key, value)
		}
		return result
	}(d.Functions)
	return fmt.Sprintf("[%v] %v functions: {%v} variables: {%v}", d.ID, d.Name, functions, variables)
}
func (d *Device) ToWebThing() webthing.WebThing {
	log.Println("[Device:toWebThing]")
	properties := func(variables map[string]Variable) map[string]webthing.Property {
		result := map[string]webthing.Property{}
		for name, variable := range variables {
			result[name] = webthing.Property{
				Type: variable.Type,
				HREF: fmt.Sprintf("/v1/devices/%v/%v", d.ID, name)}
		}
		return result
	}(d.Variables)
	wt := webthing.WebThing{
		Name:       d.Name,
		Type:       d.Type,
		Properties: properties}
	return wt
}
