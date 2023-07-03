package puppy

import "github.com/flunkey-cloud/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp("Bell Bell")
}
