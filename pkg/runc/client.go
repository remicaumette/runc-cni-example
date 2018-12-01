package runc

import (
	"github.com/pkg/errors"
	"os/exec"
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
	output, err := cmd.CombinedOutput()
	println("%s\n", string(output))
	if err != nil {
		return "", errors.Errorf("%v: %v", err.Error(), string(output))
	}
	return string(output), nil
}
