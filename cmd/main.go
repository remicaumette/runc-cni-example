package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"syscall"
)

func main() {
	exec.Command("runc", "delete", "test").Run()

	cmd := exec.Command("runc", "create", "test")
	s, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	ec := make(chan int, 1)
	go func() {
		var status int
		if err := cmd.Wait(); err != nil {
			status = 255
			if exitErr, ok := err.(*exec.ExitError); ok {
				if ws, ok := exitErr.Sys().(syscall.WaitStatus); ok {
					status = ws.ExitStatus()
				}
			}
		}
		ec <- status
		close(ec)
	}()
	fmt.Printf(" status code: %v\n", <- ec)
	r, _ := ioutil.ReadAll(s)
	println(string(r))
}
