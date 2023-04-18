package component

import (
	"context"

	"github.com/base-org/pessimism/internal/core"
	"github.com/base-org/pessimism/internal/logging"
	"go.uber.org/zap"
)

// TODO(#12): No Aggregation Component Support
type Aggregator struct {
	ctx context.Context

	*metaData
}

// NewAggregator ... Initializer
func NewAggregator(ctx context.Context, outType core.RegisterType,
	od OracleDefinition, opts ...Option) (Component, error) {
	a := &Aggregator{
		ctx: ctx,

		metaData: &metaData{
			id:             core.NilCompID(),
			cType:          core.Aggregator,
			egressHandler:  newEgressHandler(),
			ingressHandler: newIngressHandler(),
			state:          Inactive,
			output:         outType,
		},
	}

	for _, opt := range opts {
		opt(a.metaData)
	}

	if cfgErr := od.ConfigureRoutine(); cfgErr != nil {
		return nil, cfgErr
	}

	logging.WithContext(ctx).Info("Constructed component",
		zap.String("ID", a.metaData.id.String()))

	return a, nil
}
