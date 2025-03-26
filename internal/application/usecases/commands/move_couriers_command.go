package commands

type MoveCouriersCommand struct {
	isSet bool
}

func NewMoveCouriersCommand() (MoveCouriersCommand, error) {
	return MoveCouriersCommand{isSet: true}, nil
}

func (c MoveCouriersCommand) isEmpty() bool {
	return !c.isSet
}
