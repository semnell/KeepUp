package main

import (
	"fmt"
	"os"

	"github.com/semnell/KeepUp/server"
	"github.com/semnell/KeepUp/worker"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("Running server")
		server.Serve("/Users/sem/Projects/KeepUp/config.yaml")
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "worker" {
			fmt.Println("Running worker")
			worker.Work()
			fmt.Println("Worker done")
		}
	}
	server.Serve("/Users/sem/Projects/KeepUp/config.yaml")
}
