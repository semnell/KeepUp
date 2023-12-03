// Description: Entrypoint for the application. Starts the server and worker
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/semnell/KeepUp/server"
	"github.com/semnell/KeepUp/worker"
)

func main() {
	// test if .env file exists in current directory
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("No .env file found in current directory, running with env values")
	} else if err != nil {
		fmt.Printf("Error while checking for .env file: %s\n", err)
		return
	} else {
		godotenv.Load()
	}
	var confPath = os.Getenv("CONFIG_FILE_PATH")
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		fmt.Println("Config path not found defaulting to in-dir config.")
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		confPath = path + "/config.yaml"

	}
	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("Running server and worker")
		go worker.Work()
		server.Serve(confPath)
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "worker" {
			fmt.Println("Running worker")
			worker.Work()
			fmt.Println("Worker done")
		} else if os.Args[1] == "server" {
			server.Serve(confPath)
		}
	}
}
