package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// Create a sample Config instance
	config := Config{
		Version: "1.0",
		Jobs: []Job{
			{
				Name:     "TestJob",
				Scheme:   "https",
				URL:      "example.com",
				Interval: 60,
				Timeout:  10,
				Headers: []JobHeaders{
					{
						Key:   "Authorization",
						Value: "Bearer Token",
					},
				},
				Method: "GET",
				Expect: JobExpect{
					Status: 200,
					Body:   "OK",
				},
			},
		},
	}

	// Assert the fields of the Config instance
	assert.Equal(t, "1.0", config.Version)
	assert.Len(t, config.Jobs, 1)

	// Assert the fields of the Job instance within Config
	job := config.Jobs[0]
	assert.Equal(t, "TestJob", job.Name)
	assert.Equal(t, "https", job.Scheme)
	assert.Equal(t, "example.com", job.URL)
	assert.Equal(t, 60, job.Interval)
	assert.Equal(t, 10, job.Timeout)
	assert.Len(t, job.Headers, 1)
	assert.Equal(t, "Authorization", job.Headers[0].Key)
	assert.Equal(t, "Bearer Token", job.Headers[0].Value)
	assert.Equal(t, "GET", job.Method)
	assert.Equal(t, 200, job.Expect.Status)
	assert.Equal(t, "OK", job.Expect.Body)
}

func TestUpdateMetricPost(t *testing.T) {
	// Create a sample UpdateMetricPost instance
	updateMetric := UpdateMetricPost{
		URL:          "example.com",
		Reason:       "Success",
		ResCode:      200,
		MarkUp:       true,
		ResponseTime: 10.5,
	}

	// Assert the fields of the UpdateMetricPost instance
	assert.Equal(t, "example.com", updateMetric.URL)
	assert.Equal(t, "Success", updateMetric.Reason)
	assert.Equal(t, 200, updateMetric.ResCode)
	assert.True(t, updateMetric.MarkUp)
	assert.Equal(t, 10.5, updateMetric.ResponseTime)
}
