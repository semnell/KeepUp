package main

import (
	"context"
	"io"
	"log"
	"net/http"

	worker "github.com/contribsys/faktory_worker_go"
)

func processUrl(ctx context.Context, args ...interface{}) error {
	help := worker.HelperFor(ctx)
	log.Printf("Working on job %s\n", help.Jid())
	// the arg is an interface, so we need to cast it to the KeepUpConfig struct
	config := args[0].(interface{})
	configLocal := config.(Job)
	log.Printf("Processing %s\n", configLocal.Name)
	res, err := http.Get(configLocal.URL)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return err
	}
	log.Printf("Status code: %d\n", res.StatusCode)
	if res.StatusCode != configLocal.Expect.Status {
		log.Printf("Status code %d did not match expected %d\n", res.StatusCode, configLocal.Expect.Status)
		return err
	}
	if configLocal.Expect.Body != "" {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
		}
		if string(body) != configLocal.Expect.Body {
			log.Printf("Body %s did not match expected %s\n", string(body), configLocal.Expect.Body)
			return err
		}
	}
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return err
	}
	log.Printf("Status code: %d\n", res.StatusCode)
	return nil
}

func main() {
	mgr := worker.NewManager()
	mgr.Register("test", processUrl)
	mgr.Concurrency = 20

	mgr.ProcessStrictPriorityQueues("test")
	mgr.Labels = append(mgr.Labels, "test")
	mgr.Run()
}

type Job struct {
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
}
