package data

import (
	"context"

	"github.com/environment-toolkit/actuator/data/aggregates"
	"github.com/environment-toolkit/actuator/data/sagas"

	"github.com/go-apis/eventsourcing/es"
)

func NewClient(
	ctx context.Context,
	pcfg *es.ProviderConfig,
) (es.Client, error) {
	reg, err := es.NewRegistry(
		pcfg.Service,

		aggregates.NewJob(),

		sagas.NewDeploySaga(),
	)
	if err != nil {
		return nil, err
	}

	return es.NewClient(ctx, pcfg, reg)
}
