package worker

import (
	"context"
	"encoding/json"
	"net/http"

	faktoryWork "github.com/contribsys/faktory_worker_go"
	"github.com/semnell/KeepUp/utils"
)

var logger = utils.SetupSugaredLogger()

func Work() {
	mgr := faktoryWork.NewManager()
	mgr.Register("checkUrl", HandleJob)
	mgr.Concurrency = 20
	mgr.ProcessStrictPriorityQueues("urls")
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

	return nil
}

func checkUrl(job utils.Job) (err error) {
	if job.
	res, err := http.Head(job.URL)
	if err != nil {
		panic(err)
	}
	
	return nil
}
