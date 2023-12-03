package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	faktoryWork "github.com/contribsys/faktory_worker_go"
	"github.com/semnell/KeepUp/utils"
)

var logger = utils.SetupSugaredLogger()

// var res *http.Response

func Work() {
	mgr := faktoryWork.NewManager()
	mgr.Register("checkUrl", HandleJob)
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

func HandleJob(ctx context.Context, args ...interface{}) error {
	help := faktoryWork.HelperFor(ctx)
	logger.Infof("Received job: %s", help.Jid())
	obj := utils.Job{}
	err := json.Unmarshal([]byte(args[0].(string)), &obj)
	if err != nil {
		logger.Errorf("Error unmarshalling json: %v", err)
	}
	logger.Debug("running job: " + obj.Name)
	checkUrl(obj)
	return nil
}

func checkUrl(job utils.Job) (err error) {
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
		logger.Error(err.Error())
	}
	elapsed := time.Since(start)
	if err != nil {
		panic(err)
	}
	callback(job, res, elapsed)
	return nil
}

func callback(job utils.Job, res *http.Response, elapsed time.Duration) {
	var updateObj = utils.UpdateMetricPost{}
	updateObj.ResCode = res.StatusCode
	updateObj.URL = job.URL
	if updateObj.ResCode == job.Expect.Status {
		updateObj.MarkUp = true
	}
	updateObj.ResponseTime = float64(elapsed.Milliseconds())
	if job.Expect.Body != "" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		respBytes := buf.String()
		respString := string(respBytes)
		if job.Expect.Body != respString {
			updateObj.MarkUp = false
		}
	}
	b, err := json.Marshal(updateObj)
	if err != nil {
		panic(err)
	}
	request, er := http.NewRequest("POST", os.Getenv("SERVER_CALLBACK_URL"), bytes.NewBuffer(b))
	postclient := &http.Client{}
	if er != nil {
		panic(er)
	}
	response, er := postclient.Do(request)
	if er != nil {
		panic(er)
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
