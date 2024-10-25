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

// Work starts the main worker routine
func Work() {
	mgr := faktoryWork.NewManager()
	mgr.Register("checkURL", HandleJob)
	if os.Getenv("WORKER_CONCURRENCY") == "" {
		os.Setenv("WORKER_CONCURRENCY", "1")
	}
	concurrency, err := strconv.Atoi(os.Getenv("WORKER_CONCURRENCY"))
	if err != nil {
		logger.Errorf("Error converting WORKER_CONCURRENCY to int: %v", err)
		return
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
		return err
	}
	logger.Debug("running job: " + obj.Name)
	err = checkURL(obj)
	if err != nil {
		logger.Errorf("Error checking URL: %v", err)
		return err
	}
	return nil
}

// checkURL checks the status of a URL and sends a callback with the result
func checkURL(job utils.Job) error {
	if job.Scheme == "" {
		job.Scheme = "https"
	}
	var res *http.Response
	job.URL = job.Scheme + "://" + job.URL
	if job.Method == "" {
		job.Method = "HEAD"
	}
	start := time.Now()
	res, err := doRequest(job, res)
	if err != nil {
		logger.Warnf("Request error: %v", err)
		callback(job, res, time.Since(start))
		return err
	}
	elapsed := time.Since(start)
	callback(job, res, elapsed)
	return nil
}

// callback sends a callback with the result of the URL check
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
	if updateObj.ResCode == job.Expect.Status && res != nil {
		updateObj.MarkUp = true
	}
	updateObj.ResponseTime = float64(elapsed.Milliseconds())
	if job.Expect.Body != "" && res != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		respBytes := buf.String()
		respString := string(respBytes)
		if job.Expect.Body != respString {
			updateObj.MarkUp = false
		}
	}
	if job.Expect.Contains != nil && res != nil {
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
		logger.Errorf("Error marshalling updateObj: %v", err)
		return
	}
	request, localError := http.NewRequest("POST", os.Getenv("SERVER_CALLBACK_URL"), bytes.NewBuffer(b))
	if localError != nil {
		logger.Errorf("Error creating new request: %v", localError)
		return
	}
	client := &http.Client{}
	response, localError := client.Do(request)
	if localError != nil {
		logger.Errorf("Error sending request: %v", localError)
		return
	}
	defer response.Body.Close()
	logger.Infof("Ran successfully for %s", job.URL)
}

// doRequest performs the HTTP request for the given job
func doRequest(job utils.Job, res *http.Response) (*http.Response, error) {
	var err error
	if job.Method == "HEAD" {
		res, err = http.Head(job.URL)
	} else if job.Method == "GET" {
		res, err = http.Get(job.URL)
	} else {
		err = logger.Errorf("%s is not a supported method right now.", job.Method)
	}
	return res, err
}
