package main

import (
	"flag"
	"gitlab.com/expected.sh/agent/pkg/server"
	"log"
)

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	log.Printf("starting the nekomata server...")

	addr := *flag.String("addr", ":9000", "server address")
	storage := *flag.String("addr", "/var/lib/nekomata", "storage directory")
	runtime := *flag.String("addr", "/usr/local/bin/runc", "runtime path")
	cniPlugins := *flag.String("cni-plugins", "/opt/cni/plugins", "cni plugins directory")
	cniConfig := *flag.String("cni-config", "/etc/cni/net.d", "cni configuration directory")
	flag.Parse()

	log.Printf("server address: %v\n", addr)
	log.Printf("storage directory: %v\n", storage)
	log.Printf("runtime path: %v\n", runtime)
	log.Printf("cni plugins directory: %v\n", cniPlugins)
	log.Printf("cni configuration directory: %v\n", cniConfig)

	s := server.Server{
		Storage: storage,
		Runtime: runtime,
		CNIPlugins: cniPlugins,
		CNIConfig: cniConfig,
	}

	log.Printf("listening on %v...\n", addr)
	if err := s.Listen(addr); err != nil {
		log.Fatalf("an error occurred: %v\n", err)
	}
}
