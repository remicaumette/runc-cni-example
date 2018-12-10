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
		log.Output(1, fmt.Sprintf("an error occurred: %v\n", err))
		os.Exit(1)
	}
}


func setupNetworking(runtime runc.Runc, networking cni.CNI) {
	state, err := runtime.State(context.Background(), "hello")
	checkErr(err)
	log.Printf("%v\n", state.Pid)
	result, err := networking.Setup("eth0", fmt.Sprintf("/proc/%v/ns/net", state.Pid))
	checkErr(err)
	log.Printf("=== INTERFACES ===\n")
	for _, config := range result.Interfaces {
		log.Printf("sandbox: %v\n", config.Sandbox)
		for _, ip := range config.IPConfigs {
			log.Printf("ip: %v mask: %v gateway: %v\n", ip.IP.String(), ip.IP.DefaultMask(), ip.Gateway.String())
		}
	}
	log.Printf("==================\n")

}

func destroyNetworking(runtime runc.Runc, networking cni.CNI) {
	state, err := runtime.State(context.Background(), "hello")
	checkErr(err)
	networking.Remove("eth0", fmt.Sprintf("/proc/%v/ns/net", state.Pid))
}

func main() {
	runtime := runc.Runc{}
	socket, err := runc.NewTempConsoleSocket()
	checkErr(err)

	runtime.Delete(context.Background(), "hello", &runc.DeleteOpts{
		Force: true,
	})

	defer socket.Close()
	err = runtime.Create(context.Background(), "hello", "/home/vagrant/go/src/gitlab.com/expected.sh/agent", &runc.CreateOpts{
		ConsoleSocket: socket,
	})
	checkErr(err)
	err = runtime.Start(context.Background(), "hello")
	checkErr(err)
	networking, err := cni.New(cni.WithMinNetworkCount(2),
		cni.WithPluginConfDir("/etc/cni/net.d"),
		cni.WithPluginDir([]string{"/opt/cni/plugins"}),
	)
	checkErr(err)
	checkErr(networking.Load(cni.WithLoNetwork, cni.WithDefaultConf))
	setupNetworking(runtime, networking)
	defer func() {
		println()
		log.Printf("delete network...\n")
		destroyNetworking(runtime, networking)

		log.Printf("delete container...\n")
		runtime.Delete(context.Background(), "hello", &runc.DeleteOpts{
			Force: true,
		})

		log.Printf("goodbye\n")
	}()
	console, err := socket.ReceiveMaster()
	checkErr(err)
	go func() {
		io.Copy(os.Stdout, console)
	}()
	io.Copy(console, os.Stdin)
}
