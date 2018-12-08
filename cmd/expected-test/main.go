package main

import (
	"context"
	"fmt"
	"github.com/containerd/go-cni"
	"github.com/containerd/go-runc"
	"io"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Output(2, fmt.Sprintf("an error occurred: %v\n", err))
		os.Exit(1)
	}
}

func setupNetworking(runtime runc.Runc) {
	state, err := runtime.State(context.Background(), "hello")
	checkErr(err)
	log.Printf("%v\n", state.Pid)
	networking, err := cni.New(cni.WithMinNetworkCount(2),
		cni.WithPluginConfDir("/home/vagrant/go/src/gitlab.com/expected.sh/agent/net"),
		cni.WithPluginDir([]string{"/opt/cni/"}),
		 )

	checkErr(err)
	checkErr(networking.Load(cni.WithLoNetwork, cni.WithDefaultConf))
	result, err := networking.Setup("eth0", fmt.Sprintf("/proc/%v/ns/net", state.Pid), cni.WithCapabilityPortMap([]cni.PortMapping{
		{
			HostPort: 8080,
			ContainerPort: 80,
			Protocol: "tcp",
		},
	}))
	checkErr(err)
	log.Printf("%v\n", result)
}

func main() {
	runtime := runc.Runc{}
	socket, err := runc.NewTempConsoleSocket()
	checkErr(err)

	containers, err := runtime.List(context.Background())
	checkErr(err)
	for _, container := range containers {
		if container.ID == "hello" {
			checkErr(runtime.Delete(context.Background(), "hello", &runc.DeleteOpts{
				Force: true,
			}))
		}
	}

	defer socket.Close()
	err = runtime.Create(context.Background(), "hello", "/home/vagrant/go/src/gitlab.com/expected.sh/agent", &runc.CreateOpts{
		ConsoleSocket: socket,
	})
	checkErr(err)
	err = runtime.Start(context.Background(), "hello")
	checkErr(err)
	setupNetworking(runtime)
	console, err := socket.ReceiveMaster()
	checkErr(err)
	go func() {
		io.Copy(os.Stdout, console)
	}()
	io.Copy(console, os.Stdin)
}
