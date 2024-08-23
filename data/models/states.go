package models

type StateState string

const (
	StateStateNew      StateState = "new"
	StateStateUpdating StateState = "updating"
	StateStateUpdated  StateState = "updated"
	StateStateDeleting StateState = "deleting"
	StateStateDeleted  StateState = "deleted"
	StateStateFailed   StateState = "failed"
)

func (a StateState) Enum() []interface{} {
	return []interface{}{
		StateStateNew,
		StateStateUpdating,
		StateStateUpdated,
		StateStateDeleting,
		StateStateDeleted,
		StateStateFailed,
	}
}
