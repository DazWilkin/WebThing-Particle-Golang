package onion

import (
	"fmt"
	"testing"
)

var x = struct {
	username string
	password string
	apikey   string
}{
	"username",
	"password",
	"apikey"}

func TestNewAccount(t *testing.T) {
	a, _ := NewAccount(x.username, x.password, x.apikey)
	if a.Username != x.username || a.Password != x.password || a.APIKey != x.apikey {
		t.Errorf("NewAccount is incorrect, got: %v, want: %v", a, x)
	}
}
func TestString(t *testing.T) {
	a, _ := NewAccount(x.username, x.password, x.apikey)
	r := fmt.Sprintf("%v: %v (%v)", x.username, x.password, x.apikey)
	if a.String() != r {
		t.Errorf("String is incorrect, got: %v, want: %v", a.String(), r)
	}
}
