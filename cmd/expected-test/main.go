package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	r, w, err := os.Pipe()
	checkErr(err)

	cmd := exec.Command("runc", "create", "test")
	//cmd := exec.Command("sh", "-c", "echo hello world")
	cmd.Stdout = w

	in, _ := cmd.StdinPipe()
	in.Close()
	checkErr(cmd.Run())
	w.Close()

	fmt.Printf("%v\n", cmd.ProcessState.Success())
	out, err := ioutil.ReadAll(r)
	checkErr(err)
	fmt.Printf("%v", string(out))
}
