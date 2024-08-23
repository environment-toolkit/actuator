package sagas

import (
	"context"

	"github.com/environment-toolkit/actuator/data/events"
	"github.com/go-apis/eventsourcing/es"
)

type deploySaga struct {
	es.BaseSaga
}

func (s *deploySaga) HandleGridStateUpdating(ctx context.Context, evt *es.Event, data *events.GridStateUpdating) ([]es.Command, error) {
	// do nothing on everything else.
	return nil, nil
}

func NewDeploySaga() es.IsSaga {
	return &deploySaga{}
}
