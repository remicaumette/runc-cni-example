package main

import (
	"gitlab.com/expected.sh/agent/pkg/runc"
	"log"
	"os"
)

func hasContainer(client runc.Client, id string) (bool, error) {
	containers, err := client.List()
	if err != nil {
		return false, err
	}
	for _, container := range containers {
		if container.ID == id {
			return true, nil
		}
	}
	return false, nil
}

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	if len(os.Args) > 1 {
		client := runc.New("/usr/local/bin/runc")
		id := os.Args[1]

		log.Println("list")
		has, err := hasContainer(client, id)
		if err != nil {
			log.Fatalf("unable to retreive containers: %v\n", err)
		}

		if has {
			log.Println("delete")
			if err = client.Delete(id, &runc.DeleteOps{Force: true}); err != nil {
				log.Fatalf("unable to delete the container: %v\n", err)
			}
		}

		log.Println("create")
		if err = client.Create(id, &runc.CreateOpts{
			Bundle: "/home/vagrant/go/src/gitlab.com/expected.sh/agent" }); err != nil {
			log.Fatalf("unable to create the container: %v\n", err)
		}
		println("ok")
		/*	log.Println("start")
			if err = run.Start(context.Background(), id); err != nil {
				log.Fatalf("unable to start the container: %v\n", err)
			}
			log.Println("state")
			container, err := run.State(context.Background(), id)
			if err != nil {
				log.Fatalf("unable to get the container's state: %v\n", err)
			}
			log.Printf("pid: %v\n", container.Pid)
		*/
	}
}
