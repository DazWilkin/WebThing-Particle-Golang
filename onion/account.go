package onion

import (
	"fmt"
)

type Account struct {
	Username string
	Password string
	APIKey   string
}

func NewAccount(username, password, apikey string) (Account, error) {
	a := Account{
		Username: username,
		Password: password,
		APIKey:   apikey}
	return a, nil
}
func (a *Account) String() string {
	return fmt.Sprintf("%v: %v (%v)", a.Username, a.Password, a.APIKey)
}
