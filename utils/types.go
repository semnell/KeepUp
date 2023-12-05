package utils

// Config is the struct for the config file
type Config struct {
	Version string `yaml:"version"`
	Jobs    []Job  `yaml:"jobs"`
}

// Job is the struct for a job in the config file
type Job struct {
	Name     string       `yaml:"name"`
	Scheme   string       `yaml:"scheme"`
	URL      string       `yaml:"url"`
	Interval int          `yaml:"interval"`
	Timeout  int          `yaml:"timeout"`
	Headers  []JobHeaders `yaml:"headers"`
	Method   string       `yaml:"method"`
	Expect   JobExpect    `yaml:"expect"`
}

// JobHeaders is the struct for the headers in a job
type JobHeaders struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// JobExpect is the struct for the expect in a job
type JobExpect struct {
	Status   int      `yaml:"status"`
	Body     string   `yaml:"body"`
	Contains []string `yaml:"contains"`
}

// UpdateMetricPost is the struct for the callback
type UpdateMetricPost struct {
	URL          string  `json:"url"`
	Reason       string  `json:"reason"`
	ResCode      int     `json:"resCode"`
	MarkUp       bool    `json:"markUp"`
	ResponseTime float64 `json:"responseTime"`
}
