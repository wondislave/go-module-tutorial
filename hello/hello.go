package main

import (
	"fmt"

	"github.com/wondislave/go-module-tutorial/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
