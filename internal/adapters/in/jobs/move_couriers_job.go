package jobs

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
	"lisichkinuriy/delivery/internal/application/usecases/commands"
)

var _ cron.Job = &MoveCouriersJob{}

type MoveCouriersJob struct {
	moveCouriersCommandHandler commands.IMoveCouriersHandler
}

func NewMoveCouriersJob(
	moveCouriersCommandHandler commands.IMoveCouriersHandler) (*MoveCouriersJob, error) {
	if moveCouriersCommandHandler == nil {
		return nil, errors.New("moveCouriersCommandHandler")
	}

	return &MoveCouriersJob{
		moveCouriersCommandHandler: moveCouriersCommandHandler}, nil
}

func (j *MoveCouriersJob) Run() {
	ctx := context.Background()
	log.Info("move couriers job started")
	command, err := commands.NewMoveCouriersCommand()
	if err != nil {
		log.Error(err)
	}
	err = j.moveCouriersCommandHandler.Handle(ctx, command)
	if err != nil {
		log.Error(err)
	}
}
