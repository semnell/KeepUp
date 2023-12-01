package utils

type Config struct {
	Version string `yaml:"version"`
	Jobs    []Job  `yaml:"jobs"`
}

type Job struct {
	Name     string `yaml:"name"`
	Scheme   string `yaml:"scheme"`
	URL      string `yaml:"url"`
	Interval int    `yaml:"interval"`
	Timeout  int    `yaml:"timeout"`
	Headers  []struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	} `yaml:"headers"`
	Method string `yaml:"method"`
	Expect struct {
		Status int    `yaml:"status"`
		Body   string `yaml:"body"`
	} `yaml:"expect"`
}

type UpdateMetricPost struct {
	URL          string  `json:"url"`
	Reason       string  `json:"reason"`
	ResCode      int     `json:"resCode"`
	MarkUp       bool    `json:"markUp"`
	ResponseTime float64 `json:"responseTime"`
}
