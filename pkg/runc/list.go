package runc

import (
	"encoding/json"
	"time"
)

type Container struct {
	Version string `json:"ociVersion"`
	ID string `json:"id"`
	InitProcessPid int `json:"pid"`
	Status string `json:"status"`
	Bundle string `json:"bundle"`
	Rootfs string `json:"rootfs"`
	Created time.Time `json:"created"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Owner string `json:"owner"`
}

func (client *Client) List() ([]Container, error) {
	output, err := client.runCommand("list", "--format=json")
	if err != nil {
		return nil, err
	}
	var containers []Container
	if err = json.Unmarshal([]byte(output), &containers); err != nil {
		return nil, err
	}
	return containers, nil
}
