package main

import (
	"fmt"
	"net/http"

	"github.com/RaphSku/Go-Full-Template/exporters"
	"github.com/RaphSku/Go-Full-Template/handlers"
	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	logger := hclog.Default()
	pe := exporters.GetPrometheusExporter(logger)
	metricProvider := metric.NewMeterProvider(metric.WithReader(pe))

	bh := handlers.NewBaseHandler(logger, metricProvider)

	prometheushost := "localhost"
	prometheusport := ":2223"
	logger.Info(fmt.Sprintf("serving metrics at %s%s/metrics", prometheushost, prometheusport))
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(fmt.Sprintf("%s%s", prometheushost, prometheusport), nil)
		if err != nil {
			logger.Error("error serving http: %v", err)
			return
		}
	}()

	echoserver := echo.New()

	echoserver.Use(middleware.Logger())
	echoserver.Use(middleware.Recover())

	echoserver.GET("/", bh.GetBase)

	echoport := ":1323"
	echohost := "localhost"
	logger.Info(fmt.Sprintf("echoserver started serving at %s%s", echohost, echoport))
	echoserver.Logger.Fatal(echoserver.Start(fmt.Sprintf("%s%s", echohost, echoport)))
}
