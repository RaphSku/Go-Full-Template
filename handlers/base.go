package handlers

import (
	"github.com/hashicorp/go-hclog"
	"go.opentelemetry.io/otel/sdk/metric"
)

type BaseHandler struct {
	logger hclog.Logger
	mp     *metric.MeterProvider
}

func NewBaseHandler(logger hclog.Logger, mp *metric.MeterProvider) *BaseHandler {
	return &BaseHandler{logger, mp}
}
