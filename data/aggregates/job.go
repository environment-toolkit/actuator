package aggregates

import (
	"context"

	"github.com/environment-toolkit/actuator/data/commands"
	"github.com/environment-toolkit/actuator/data/events"

	"github.com/go-apis/eventsourcing/es"
)

type Job struct {
	es.BaseAggregateSourced
}

func (a *Job) HandleNewEnvironment(ctx context.Context, cmd *commands.NewJob) error {
	return a.Apply(ctx, &events.JobCreated{})
}

func NewJob() *Job {
	return &Job{}
}
