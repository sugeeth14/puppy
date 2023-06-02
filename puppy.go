package puppy

import "github.com/sugeeth14/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!!!"
}

func BarkLoudly() string {
	return dog.WhenGrownUp("Loud BARK!!!!")
}
