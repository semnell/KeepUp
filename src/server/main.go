package main

import (
	"fmt"
	"log"
	"os"
	"time"

	faktory "github.com/contribsys/faktory/client"
	"gopkg.in/yaml.v2"
)

var err error

type KeepUpConfig struct {
	Version string `yaml:"version"`
	Jobs    []struct {
		Name     string `yaml:"name"`
		Type     string `yaml:"type,omitempty"`
		URL      string `yaml:"url"`
		Interval int    `yaml:"interval,omitempty"`
		Timeout  int    `yaml:"timeout,omitempty"`
		Headers  []struct {
			Key   string `yaml:"key"`
			Value string `yaml:"value"`
		} `yaml:"headers,omitempty"`
		Method string `yaml:"method,omitempty"`
		Expect struct {
			Status int    `yaml:"status"`
			Body   string `yaml:"body,omitempty"`
		} `yaml:"expect,omitempty"`
	} `yaml:"jobs"`
}

func main() {
	localConfig := KeepUpConfig{}

	// load config into KeepUpConfig
	// if KEEPUP_CONFIG is set, load that file else load config.yaml in current directory
	if err != nil {
		log.Fatalf("error connecting to faktory: %v", err)
	}
	var path string

	if os.Getenv("KEEPUP_CONFIG") != "" {
		fmt.Println("loading config from KEEPUP_CONFIG")
		path = os.Getenv("KEEPUP_CONFIG")

	} else {
		fmt.Println("loading config from current directory")
		// get current dir
		path, err = os.Getwd()
		path = path + "/config.yaml"
		if err != nil {
			log.Fatalf("error getting current directory: %v", err)
		}
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0600)

	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	dec := yaml.NewDecoder(file)
	err = dec.Decode(&localConfig)
	if err != nil {
		err = fmt.Errorf("error loading config file %v!", err)
		panic(err)
	} else {
		fmt.Println("config loaded successfully!")
	}
	if len(localConfig.Jobs) == 0 {
		err = fmt.Errorf("no jobs found in config file!")
		panic(err)
	}

	for _, job := range localConfig.Jobs {
		if job.Interval > 0 {
			go postJob(job.Interval, job)
		} else {
			go postJob(15, job)
		}
	}
	for {
		time.Sleep(time.Second * 5)
	}
}

func postJob(interval int, obj interface{}) {
	client, err := faktory.Open()
	for range time.Tick(time.Second * time.Duration(interval)) {
		fmt.Println("posting job to faktory")
		job := faktory.NewJob("test", obj)
		if err != nil {
			panic(err)
		}
		job.Queue = "test"
		err = client.Push(job)
		if err != nil {
			panic(err)
		}

	}
}
