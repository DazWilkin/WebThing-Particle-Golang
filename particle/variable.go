package particle

import (
	"fmt"
	"log"
)

type Variable struct {
	Type  string
	Value string
}

func (m *Variable) String() string {
	log.Println("[Variable:String]")
	return fmt.Sprintf("%v(%v)", m.Value, m.Type)
}
