package events

import (
	"time"

	"github.com/environment-toolkit/actuator/data/models"
	"github.com/go-apis/eventsourcing/es"

	"github.com/google/uuid"
)

type GridStateUpdating struct {
	es.BaseEvent `es:"StateUpdating;service=grid"`

	State      models.StateState `json:"state"`
	SpecId     uuid.UUID         `json:"spec_id"`
	Target     models.Target     `json:"target"`
	UpdatingAt time.Time         `json:"updating_at"`
}
