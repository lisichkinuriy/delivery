package commands

type AssignOrderCommand struct {
	isSet bool
}

func NewAssignOrderCommand() (AssignOrderCommand, error) {
	return AssignOrderCommand{isSet: true}, nil
}

func (c AssignOrderCommand) isEmpty() bool {
	return !c.isSet
}
