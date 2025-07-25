package prometheus

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Handler - display Prometheus metrics
func Handler() func(c *gin.Context) {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		if !conf.AppConfig.PrometheusEnable {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

var up = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "watch_your_lan",
	Name:      "up",
	Help:      "Whether the host is up (1 for yes, 0 for no)",
}, []string{"ip", "iface", "name", "mac", "known"})

// Add a Prometheus metric
func Add(oneHist models.Host) {
	if oneHist.Name == "" {
		oneHist.Name = "unknown"
	}

	up.With(prometheus.Labels{
		"ip":    oneHist.IP,
		"iface": oneHist.Iface,
		"name":  oneHist.Name,
		"mac":   oneHist.Mac,
		"known": strconv.Itoa(oneHist.Known),
	}).Set(float64(oneHist.Now))
}
