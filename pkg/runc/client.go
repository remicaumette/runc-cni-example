package runc

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os/exec"
	"syscall"
)

type Client struct {
	Command	string
}

func New(command string) Client {
	return Client{
		Command: command,
	}
}

func (client *Client) runCommand(args ...string) (string, error) {
	cmd := exec.Command(client.Command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}
	statusCode := make(chan int, 1)
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
		statusCode <- status
		close(statusCode)
	}()
	if <- statusCode == 0 {
		output, _ := ioutil.ReadAll(stdout)
		if err != nil {
			return "", err
		}
		return string(output), nil
	}
	output, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}
	return "", errors.New(string(output))
}
