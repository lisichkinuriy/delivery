package jobs

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
	"lisichkinuriy/delivery/internal/application/usecases/commands"
)

var _ cron.Job = &AssignOrdersJob{}

type AssignOrdersJob struct {
	assignOrdersCommandHandler commands.IAssignOrderHandler
}

func NewAssignOrdersJob(
	assignOrdersCommandHandler commands.IAssignOrderHandler) (*AssignOrdersJob, error) {
	if assignOrdersCommandHandler == nil {
		return nil, errors.New("moveCouriersCommandHandler")
	}

	return &AssignOrdersJob{
		assignOrdersCommandHandler: assignOrdersCommandHandler}, nil
}

func (j *AssignOrdersJob) Run() {
	ctx := context.Background()
	log.Info("assign orders job started")
	command, err := commands.NewAssignOrderCommand()
	if err != nil {
		log.Error(err)
	}
	err = j.assignOrdersCommandHandler.Handle(ctx, command)
	if err != nil {
		log.Error(err)
	}
}
