package exporters

import (
	"github.com/hashicorp/go-hclog"
	"go.opentelemetry.io/otel/exporters/prometheus"
)

func GetPrometheusExporter(logger hclog.Logger) *prometheus.Exporter {
	exporter, err := prometheus.New()
	if err != nil {
		logger.Error(err.Error())
	}

	return exporter
}
