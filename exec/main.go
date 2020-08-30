package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//cmd := exec.Command("(Get-Host).UI.RawUI.WindowSize")
	var out bytes.Buffer
	cmd := exec.Command("powershell","/c","(Get-Host).UI.RawUI.WindowSize")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
	//cmd.Run()
	//fmt.Println(h,w)
}
