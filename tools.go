package cli

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func mkClsFuncs() map[string]func() {
	clsFuncs := make(map[string]func())
	clsFuncs["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clsFuncs["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	return clsFuncs
}

func ClearScreen() {
	clsFuncs := mkClsFuncs()
	clsFunc, ok := clsFuncs[runtime.GOOS]
	if ok {
		clsFunc()
	} else {
		log.Fatal("unknown platform.")
	}
}
