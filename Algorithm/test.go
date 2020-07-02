package main

import (
	. "fmt"
	"os/exec"
)

func main() {
	app := "g++"
	arg0 := "ValidParentheses.c++"
	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		Println(err.Error())
		return
	}

	Print(string(stdout))
}
