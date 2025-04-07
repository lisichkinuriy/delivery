package order

const (
	StatusCreated   Status = "Created"
	StatusAssigned  Status = "Assigned"
	StatusCompleted Status = "Completed"
)

type Status string

func (s Status) Equals(other Status) bool { return s == other }

func (s Status) String() string {
	return string(s)
}
