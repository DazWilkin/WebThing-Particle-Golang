package particle

import (
	"fmt"
	"log"
)

type Account struct {
	Email   string
	Devices []Device
}

func NewAccount(email string, devices []Device) (Account, error) {
	log.Println("[NewAccount] Entered")
	a := Account{Email: email, Devices: devices}
	return a, nil
}
func (a *Account) AddDevice(device Device) {
	a.Devices = append(a.Devices, device)
}
func (a *Account) String() string {
	devices := func(devices []Device) string {
		log.Println("[Account:x] Entered")
		var result string
		for _, device := range devices {
			result += fmt.Sprintf("%v, ", device.ID)
		}
		return result
	}(a.Devices)
	return fmt.Sprintf("%v {%v}", a.Email, devices)
}
