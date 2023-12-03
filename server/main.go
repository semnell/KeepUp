// Package server Handles all server-side logic
package server

import (
	"io"
	"os"
	"strings"
	"time"

	"encoding/json"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/semnell/KeepUp/utils"
)

var server = gin.Default()
var logger = utils.SetupLogger()

// Serve starts the server
func Serve(confPath string) {
	// load config into utils.Config struct
	conf := utils.LoadConfig(confPath)
	registerRoutes(server)
	initPrometheusMetrics(conf)
	server.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	server.Use(ginzap.RecoveryWithZap(logger, true))
	utils.RegisterWorkers(conf)
	server.Run(":" + os.Getenv("SERVER_PORT"))
}

func registerRoutes(s *gin.Engine) (err error) {
	logger.Info("Registering routes")
	s.GET("/metrics", prometheusHandler())
	s.POST("/callback", upMarkerHandler())
	return nil
}

func initPrometheusMetrics(conf utils.Config) (err error) {
	logger.Info("Registering prometheus metrics")
	for _, job := range conf.Jobs {
		siteUpGauge.WithLabelValues(job.URL, "200").Set(0)
		siteResponseTimeGauge.WithLabelValues(job.URL).Set(0)
		logger.Info("Registered prometheus metric: " + job.Name)
	}
	prometheus.MustRegister(siteUpGauge, siteResponseTimeGauge)
	return nil
}

func prometheusHandler() gin.HandlerFunc {
	logger.Info("Registering prometheus handler")
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func upMarkerHandler() gin.HandlerFunc {
	logger.Info("Registering up marker handler")
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error(err.Error())
		}
		// load json into utils.UpdateMetricPost struct
		var content utils.UpdateMetricPost
		err = json.Unmarshal(jsonData, &content)
		if err != nil {
			logger.Error(err.Error())
		}
		// set prometheus metric
		// content.URL contains scheme, remove it
		content.URL = strings.Split(content.URL, "://")[1]
		if content.MarkUp {
			siteUpGauge.WithLabelValues(content.URL, "200").Set(1)
		} else {
			siteUpGauge.WithLabelValues(content.URL, "200").Set(0)
		}
		siteResponseTimeGauge.WithLabelValues(content.URL).Set(content.ResponseTime)

	}
}

var siteUpGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "keepup",
		Subsystem: "uptime",
		Name:      "site_up",
		Help:      "a gauge of whether the site is up in the most recent check",
	},
	[]string{
		// Which user has requested the operation?
		"url",
		// Of what type is the operation?
		"response_code",
	},
)

var siteResponseTimeGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "keepup",
		Subsystem: "response_time",
		Name:      "site_response_time_in_ms",
		Help:      "a gauge of the response time of the site in the most recent check",
	},
	[]string{
		"url",
	},
)
