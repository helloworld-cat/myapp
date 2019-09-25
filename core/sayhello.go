package core

import (
	"fmt"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %s !", name)
}

func SayHi() string {
	return "Hi!"
}
