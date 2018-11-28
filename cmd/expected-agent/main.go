package main

import (
	"context"
	"fmt"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("an error occurred: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	client, err := containerd.New("/run/containerd/containerd.sock")
	checkError(err)
	defer client.Close()
	fmt.Printf("connected to the client\n")
	ctx := namespaces.WithNamespace(context.Background(), "default")

	//fmt.Printf("pulling redis image...\n")
	//image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
	image, err := client.GetImage(credis:alpine")
	checkError(err)
	//fmt.Printf("pulled (%v)!\n", image.Target().Digest.String())

	fmt.Printf("creating redis container...\n")
	container, err := client.NewContainer(ctx, "redis",
		containerd.WithNewSnapshot("redis-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)))
	checkError(err)
	defer container.Delete(ctx, containerd.WithSnapshotCleanup)
	fmt.Printf("created (%v)!\n", container.ID())

	fmt.Printf("starting the container...\n")
	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	checkError(err)
	defer task.Delete(ctx)
	statusCodeChannel, err := task.Wait(ctx)
	checkError(err)
	checkError(task.Start(ctx))

	statusCode := <- statusCodeChannel
	fmt.Printf("process exited with status: %v\n", statusCode.ExitCode())
}
