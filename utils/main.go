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

func SetupLogger() (logger *zap.Logger) {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer l.Sync()
	return l
}

func SetupSugaredLogger() (sugaredLogger *zap.SugaredLogger) {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := l.Sugar()
	defer l.Sync()
	return sugar
}

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

func RegisterWorkers(conf Config) {
	logger := SetupLogger()
	logger.Info("Registering workers")
	for _, job := range conf.Jobs {
		logger.Info("Registering worker: " + job.Name)
		if job.Interval == 0 {
			job.Interval = 60
		}
		// convert job to json
		jobJson, err := json.Marshal(job)
		if err != nil {
			panic(err)
		}

		go postJob(job.Interval, string(jobJson))
	}
	logger.Info("Workers registered")
}

func postJob(interval int, obj string) {
	client, err := faktory.Open()
	if err != nil {
		panic(err)
	}
	logger := SetupLogger()
	logger.Info(fmt.Sprintf("Posting job every %d seconds", interval))
	for range time.Tick(time.Second * time.Duration(interval)) {
		job := faktory.NewJob("checkUrl", obj)
		job.Queue = "urls"
		err = client.Push(job)
		if err != nil {
			panic(err)
		}
	}
}
