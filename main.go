package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/semnell/KeepUp/server"
	"github.com/semnell/KeepUp/worker"
)

func main() {
	// test if .env file exists in current directory
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("No .env file found in current directory, running with default values")
	} else {
		godotenv.Load()
	}

	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("Running server and worker")
		go worker.Work()
		server.Serve(os.Getenv("CONFIG_FILE_PATH"))
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "worker" {
			fmt.Println("Running worker")
			worker.Work()
			fmt.Println("Worker done")
		} else if os.Args[1] == "server" {
			server.Serve(os.Getenv("CONFIG_FILE_PATH"))
		}
	}

}
