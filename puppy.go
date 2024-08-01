package puppy

import (
	"fmt"
	"strings"
)

func Bark() {
	fmt.Println("Woof!")
}
func Barks() {
	fmt.Println("Woof! Woof! Woof!")
}
func BarkUpper() {
	fmt.Println(strings.ToUpper("Woof! Woof! Woof!"))
}
