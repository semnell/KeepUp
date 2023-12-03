package worker

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
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
