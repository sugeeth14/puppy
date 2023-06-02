package puppy

import (
	"fmt"

	"github.com/sugeeth14/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!!!"
}

func BarkLoudly() string {
	return dog.WhenGrownUp("Loud BARK!!!!")
}

func FromVersion0() {
	fmt.Println("Hello I am from version 0")
}
