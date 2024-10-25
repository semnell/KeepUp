package worker

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/semnell/KeepUp/utils"
	"github.com/stretchr/testify/assert"
)

func TestWork(t *testing.T) {
	// Set up environment variables for testing
	os.Setenv("WORKER_CONCURRENCY", "1")
	os.Setenv("JOB_QUEUE_NAME", "testQueue")

	// Capture the standard output for testing
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the Work function
	go Work()
	time.Sleep(3000)
	// Restore the standard output
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.Contains(t, string(out), "")
}

func TestCheckUrl(t *testing.T) {
	// Create a sample job object
	job := utils.Job{
		Name:   "testJob",
		URL:    "example.com",
		Scheme: "https",
		Method: "GET",
		// Add other fields as needed
		Expect: utils.JobExpect{
			Status: http.StatusOK,
			Body:   "testBody",
		},
	}

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Add mock response body if needed
	}))
	defer server.Close()

	// Set the mock server URL in the environment variable
	os.Setenv("SERVER_CALLBACK_URL", server.URL)

	// Call the checkUrl function
	err := checkURL(job)
	assert.NoError(t, err)
}

func TestHandleJob(t *testing.T) {
	// Create a sample job object
	job := utils.Job{
		Name:   "testJob",
		URL:    "example.com",
		Scheme: "https",
		Method: "GET",
		Expect: utils.JobExpect{
			Status: http.StatusOK,
			Body:   "testBody",
		},
	}

	// Convert job object to JSON string
	jobJSON, err := json.Marshal(job)
	assert.NoError(t, err)

	// Call the HandleJob function
	err = HandleJob(context.Background(), string(jobJSON))
	assert.NoError(t, err)
}

func TestDoRequest(t *testing.T) {
	// Create a sample job object
	job := utils.Job{
		Name:   "testJob",
		URL:    "example.com",
		Scheme: "https",
		Method: "GET",
	}

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Add mock response body if needed
	}))
	defer server.Close()

	// Set the mock server URL in the job object
	job.URL = server.URL

	// Call the doRequest function
	res, err := doRequest(job, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCallback(t *testing.T) {
	// Create a sample job object
	job := utils.Job{
		Name:   "testJob",
		URL:    "example.com",
		Scheme: "https",
		Method: "GET",
		Expect: utils.JobExpect{
			Status: http.StatusOK,
			Body:   "testBody",
		},
	}

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Add mock response body if needed
	}))
	defer server.Close()

	// Set the mock server URL in the environment variable
	os.Setenv("SERVER_CALLBACK_URL", server.URL)

	// Create a sample response object
	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("testBody")),
	}

	// Call the callback function
	callback(job, res, time.Millisecond*100)
}
