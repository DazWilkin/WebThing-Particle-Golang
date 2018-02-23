package particle

import (
	"fmt"
	"log"
	"time"

	"github.com/DazWilkin/webthing/webthing"
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
	p := Device{
		ID:        id,
		Name:      name,
		Type:      _type,
		Firmware:  firmware,
		Handshake: handshake,
		Functions: functions,
		Variables: variables}
	return p, nil
}
func (p *Device) AddFunction(name string, function Function) {
	log.Println("[AddFunction] Entered")
	p.Functions[name] = function
}
func (p *Device) AddVariable(name string, variable Variable) {
	log.Println("[AddVariable] Entered")
	p.Variables[name] = variable
}
func (p *Device) String() string {
	log.Println("[Device:String]")
	variables := func(variables map[string]Variable) string {
		log.Println("[Device:String:λVariables] Entered")
		result := ""
		for key, value := range variables {
			log.Printf("%v=%v\n", key, value)
			result += fmt.Sprintf("%v: %v,", key, value)
		}
		return result
	}(p.Variables)
	functions := func(functions map[string]Function) string {
		log.Println("[Device:String:λFunctions] Entered")
		result := ""
		for key, value := range functions {
			log.Printf("%v=%v\n", key, value)
			result += fmt.Sprintf("%v: %v", key, value)
		}
		return result
	}(p.Functions)
	return fmt.Sprintf("[%v] %v functions: {%v} variables: {%v}", p.ID, p.Name, functions, variables)
}
func (p *Device) ToWebThing() webthing.WebThing {
	log.Println("[Device:toWebThing]")
	properties := func(variables map[string]Variable) map[string]webthing.Property {
		result := map[string]webthing.Property{}
		for name, variable := range variables {
			result["name"] = webthing.Property{
				Type: variable.Type,
				HREF: fmt.Sprintf("/v1/devices/%v/%v", p.ID, name)}
		}
		return result
	}(p.Variables)
	wt := webthing.WebThing{
		Name:       p.Name,
		Type:       p.Type,
		Properties: properties}
	return wt
}
