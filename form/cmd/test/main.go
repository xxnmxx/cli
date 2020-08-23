package main

import (
	//"fmt"
	"os"

	"github.com/xxnmxx/cli/form"
)

func main() {
	f := form.NewForm(os.Stdin, os.Stdout, "> ")
	f.CreateList("greet", "hello", "hi", "gm", "greetings")
	f.Input("greet","float")
	f.CreateList("yep","yes","yeees","y")
	f.Input("yep","string")
}
