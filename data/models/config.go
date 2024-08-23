package models

import "github.com/google/uuid"

type Target struct {
	EnvironmentId uuid.UUID `json:"environment_id"`
	Region        string    `json:"region"`
}
