package courier

const (
	StatusFree Status = "free"
	StatusBusy Status = "busy"
)

type Status string

func (s Status) Equals(other Status) bool { return s == other }
