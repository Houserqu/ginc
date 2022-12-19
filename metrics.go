package ginc

import "github.com/penglongli/gin-metrics/ginmetrics"

var Metrics *ginmetrics.Monitor

func init() {
	Metrics = ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	Metrics.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	Metrics.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	Metrics.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
}
