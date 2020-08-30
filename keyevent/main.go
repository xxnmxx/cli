// Tooooooo slow! Can not use this!
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	input()
}

func input() {
	vk := make([]string,4)
	for {
		vk = getKeyEvent()
		if vk[0] == "40" {
			fmt.Println("under arrow.\ngood bye\n")
			break
		}
	}
}

func getKeyEvent() []string {
	cmd := exec.Command("powershell", "/c", "(get-host).ui.rawui.readkey().tostring()")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	ret := strings.Split(string(out), ",")
	return ret
}
