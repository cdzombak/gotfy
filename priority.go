package gotfy

// Priority is an enum for Ntfy's message priorities.
// See: https://docs.ntfy.sh/publish/#message-priority
type Priority int8

const (
	PriorityUnspecified = Priority(0)
	PriorityMin         = Priority(1)
	PriorityLow         = Priority(2)
	PriorityDefault     = Priority(3)
	PriorityHigh        = Priority(4)
	PriorityMax         = Priority(5)
	PriorityUrgent      = PriorityMax
)
