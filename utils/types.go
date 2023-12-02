package utils

type Config struct {
	Version string `yaml:"version"`
	Jobs    []Job  `yaml:"jobs"`
}

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

type JobHeaders struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type JobExpect struct {
	Status int    `yaml:"status"`
	Body   string `yaml:"body"`
}

type UpdateMetricPost struct {
	URL          string  `json:"url"`
	Reason       string  `json:"reason"`
	ResCode      int     `json:"resCode"`
	MarkUp       bool    `json:"markUp"`
	ResponseTime float64 `json:"responseTime"`
}
