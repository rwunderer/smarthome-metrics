package controllers

import (
	"context"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type SmarthomeAppliance interface {
	Run(ctx context.Context, metrics *metric.Metrics) error
	SetValue(ctx context.Context, fieldName string, desiredValue float64) error
	Close(ctx context.Context)
}

type SmarthomeController interface {
	Reconcile(ctx context.Context, metrics metric.MetricsMap)
}

type ApplianceMap map[string]SmarthomeAppliance
type ControllersMap map[string]SmarthomeController
