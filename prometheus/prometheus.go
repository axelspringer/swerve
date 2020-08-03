package phm

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusHTTPMetric struct {
	ClientConnections     prometheus.Gauge
	RedirectRequestsTotal *prometheus.CounterVec
	ResponseTimeHistogram *prometheus.HistogramVec
}

func NewPHM() *PrometheusHTTPMetric {
	phm := PrometheusHTTPMetric{
		ClientConnections: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "active_client_connections",
			Help: "Number of active client connections",
		}),
		RedirectRequestsTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "redirect_requests_total",
			Help: "total HTTP requests processed",
		}, []string{"code", "type"},
		),
		ResponseTimeHistogram: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "response_time",
			Help:    "Histogram of response time for handler",
			Buckets: prometheus.LinearBuckets(0, 0.25, 12),
		}, []string{"type", "code"}),
	}

	return &phm
}
func (phm *PrometheusHTTPMetric) WrapHandler(typeLabel string, handler http.Handler) http.Handler {
	wrappedHandler := promhttp.InstrumentHandlerInFlight(phm.ClientConnections,
		promhttp.InstrumentHandlerCounter(phm.RedirectRequestsTotal.MustCurryWith(prometheus.Labels{"type": typeLabel}),
			promhttp.InstrumentHandlerDuration(phm.ResponseTimeHistogram.MustCurryWith(prometheus.Labels{"type": typeLabel}),
				handler),
		),
	)
	return wrappedHandler
}
