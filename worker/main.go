// Package worker This file contains the worker code that is responsible for handeling all jobs
package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	faktoryWork "github.com/contribsys/faktory_worker_go"
	"github.com/semnell/KeepUp/utils"
)

var logger = utils.SetupSugaredLogger()

// var res *http.Response

// Work starts the main worker routine
func Work() {
	mgr := faktoryWork.NewManager()
	mgr.Register("checkURL", HandleJob)
	if os.Getenv("WORKER_CONCURRENCY") == "" {
		os.Setenv("WORKER_CONCURRENCY", "1")
	}
	concurrency, err := strconv.Atoi(os.Getenv("WORKER_CONCURRENCY"))
	if err != nil {
		panic(err)
	}
	mgr.Concurrency = concurrency
	mgr.ProcessStrictPriorityQueues(os.Getenv("JOB_QUEUE_NAME"))
	mgr.Run()
}

// HandleJob is the function that handles the job
func HandleJob(ctx context.Context, args ...interface{}) error {
	help := faktoryWork.HelperFor(ctx)
	logger.Infof("Received job: %s", help.Jid())
	obj := utils.Job{}
	err := json.Unmarshal([]byte(args[0].(string)), &obj)
	if err != nil {
		logger.Errorf("Error unmarshalling json: %v", err)
	}
	logger.Debug("running job: " + obj.Name)
	checkURL(obj)
	return nil
}

func checkURL(job utils.Job) (err error) {
	if job.Scheme == "" {
		job.Scheme = "https"
	}
	var res *http.Response
	job.URL = job.Scheme + "://" + job.URL
	if job.Method == "" {
		job.Method = "HEAD"
	}
	start := time.Now()
	res, err = doRequest(job, res, err)
	if err != nil {
		logger.Warn(err.Error())
		callback(job, res, time.Since(start))
	}
	elapsed := time.Since(start)
	if err != nil {
		logger.Error(err.Error())
	}
	callback(job, res, elapsed)
	return nil
}

func callback(job utils.Job, res *http.Response, elapsed time.Duration) {
	var updateObj = utils.UpdateMetricPost{}
	updateObj.MarkUp = false // default to false
	// test if res contains anything
	if res != nil {
		updateObj.ResCode = res.StatusCode
	} else {
		updateObj.ResCode = 0
		logger.Error("setting rescode to 0 to reflect connection error, check logs/url")
	}
	updateObj.URL = job.URL
	if updateObj.ResCode == job.Expect.Status && res != nil  {
		updateObj.MarkUp = true
	}
	updateObj.ResponseTime = float64(elapsed.Milliseconds())
	if job.Expect.Body != "" && res != nil{
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		respBytes := buf.String()
		respString := string(respBytes)
		if job.Expect.Body != respString {
			updateObj.MarkUp = false
		}
	}
	if job.Expect.Contains != nil && res != nil{
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		respBytes := buf.String()
		respString := string(respBytes)
		for _, s := range job.Expect.Contains {
			if !strings.Contains(respString, s) {
				updateObj.MarkUp = false
			}
		}

	}
	b, err := json.Marshal(updateObj)
	if err != nil {
		panic(err)
	}
	request, localError := http.NewRequest("POST", os.Getenv("SERVER_CALLBACK_URL"), bytes.NewBuffer(b))
	client := &http.Client{}
	if localError != nil {
		panic(localError)
	}
	response, localError := client.Do(request)
	if localError != nil {
		panic(localError)
	}
	defer response.Body.Close()
	logger.Info("Ran successfully for " + job.URL)
}

func doRequest(job utils.Job, res *http.Response, err error) (*http.Response, error) {
	if job.Method == "HEAD" {
		res, err = http.Head(job.URL)
	} else if job.Method == "GET" {
		res, err = http.Get(job.URL)
	} else {
		logger.Error(job.Method + " is not a supporter Method right now.")
	}
	return res, err
}
