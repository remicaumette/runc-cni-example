package runc

import (
	"encoding/json"
	"os/exec"
	"time"
)

// Container represents the platform agnostic pieces relating to a
// running container's status and state
type Container struct {
	// Version is the OCI version for the container
	Version string `json:"ociVersion"`
	// ID is the container ID
	ID string `json:"id"`
	// InitProcessPid is the init process id in the parent namespace
	InitProcessPid int `json:"pid"`
	// Status is the current status of the container, running, paused, ...
	Status string `json:"status"`
	// Bundle is the path on the filesystem to the bundle
	Bundle string `json:"bundle"`
	// Rootfs is a path to a directory containing the container's root filesystem.
	Rootfs string `json:"rootfs"`
	// Created is the unix timestamp for the creation time of the container in UTC
	Created time.Time `json:"created"`
	// Annotations is the user defined annotations added to the config.
	Annotations map[string]string `json:"annotations,omitempty"`
	// The owner of the state directory (the owner of the container).
	Owner string `json:"owner"`
}

// List returns all containers
func (client *Client) List() ([]Container, error) {
	cmd := exec.Command(client.Command, "list", "--format=json")
	bytes, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var containers []Container
	if err = json.Unmarshal(bytes, &containers); err != nil {
		return nil, err
	}
	return containers, nil
}
