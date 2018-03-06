package onion

import "testing"

var o = struct {
	name     string
	mac      string
	password string
	cloudid  string
}{
	"Omega-38A5",
	"40A36BC038A5",
	"password",
	"7435c16e-7bd5-40e2-a066-9d6529545639"}

func TestNewDevice(t *testing.T) {
	d, _ := NewDevice(o.name, o.mac, o.password, o.cloudid)
	if d.Name != o.name || d.MAC != o.mac || d.Password != o.password || d.CloudID != o.cloudid {
		t.Errorf("NewDevice is incorrect, got: %v, want: %v", d, o)
	}

}

func TestDeviceString(t *testing.T) {}
func TestToWebThing(t *testing.T)   {}
