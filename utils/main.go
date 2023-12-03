// Package utils contains helper functions and structs
package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	faktory "github.com/contribsys/faktory/client"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// SetupLogger sets up a zap logger with default settings
func SetupLogger() (logger *zap.Logger) {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer l.Sync()
	return l
}

// SetupSugaredLogger sets up a zap sugared logger with default settings
func SetupSugaredLogger() (sugaredLogger *zap.SugaredLogger) {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := l.Sugar()
	defer l.Sync()
	return sugar
}

// LoadConfig loads a yaml config file into a Config struct
func LoadConfig(path string) (conf Config) {
	logger := SetupLogger()
	logger.Info("Loading config")
	conf = Config{}
	// load yaml file into conf struct
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	logger.Info("Config loaded")
	return conf
}

// RegisterWorkers registers all jobs in the config file as workers
func RegisterWorkers(conf Config) {
	logger := SetupLogger()
	logger.Info("Registering workers")
	for _, job := range conf.Jobs {
		logger.Info("Registering worker: " + job.Name)
		if job.Interval == 0 {
			job.Interval = 60
		}
		// convert job to json
		jobJSON, err := json.Marshal(job)
		if err != nil {
			panic(err)
		}

		go postJob(job.Interval, string(jobJSON))
	}
	logger.Info("Workers registered")
}

// postJob posts a job to the faktory queue every interval seconds
func postJob(interval int, obj string) {
	client, err := faktory.Open()
	logger := SetupLogger()
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info(fmt.Sprintf("Posting job every %d seconds", interval))
	for range time.Tick(time.Second * time.Duration(interval)) {
		job := faktory.NewJob("checkUrl", obj)
		job.Queue = os.Getenv("JOB_QUEUE_NAME")
		err = client.Push(job)
		if err != nil {
			panic(err)
		}
	}
}
