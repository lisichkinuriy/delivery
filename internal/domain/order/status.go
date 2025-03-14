package order

const (
	StatusCreated   Status = "created"
	StatusAssigned  Status = "assigned"
	StatusCompleted Status = "completed"
)

type Status string

func (s Status) Equals(other Status) bool { return s == other }
