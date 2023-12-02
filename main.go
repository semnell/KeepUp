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
	var conf_path string = os.Getenv("CONFIG_FILE_PATH")
	if _, err := os.Stat(conf_path); os.IsNotExist(err) {
		fmt.Println("Config path not found defaulting to in-dir config.")
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		conf_path = path + "/config.yaml"

	}
	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("Running server and worker")
		go worker.Work()
		server.Serve(conf_path)
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "worker" {
			fmt.Println("Running worker")
			worker.Work()
			fmt.Println("Worker done")
		} else if os.Args[1] == "server" {
			server.Serve(conf_path)
		}
	}
}
