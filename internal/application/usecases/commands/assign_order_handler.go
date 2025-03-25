package commands

import (
	"context"
	"errors"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/services"
)

type IAssignOrderHandler interface {
	Handle(context.Context, AssignOrderCommand) error
}

var _ IAssignOrderHandler = &AssignOrderHandler{}

type AssignOrderHandler struct {
	unitOfWork      ports.IUnitOfWork
	courierRepo     ports.ICourierRepository
	orderRepo       ports.IOrderRepository
	orderDispatcher services.IOrderDispatcher
}

func NewAssignOrderHandler(
	unitOfWork ports.IUnitOfWork,
	courierRepo ports.ICourierRepository,
	orderRepo ports.IOrderRepository,
	orderDispatcher services.IOrderDispatcher,
) (*AssignOrderHandler, error) {

	if unitOfWork == nil {
		return nil, errors.New("unitOfWork is nil")
	}
	if courierRepo == nil {
		return nil, errors.New("courierRepo is nil")
	}
	if orderRepo == nil {
		return nil, errors.New("orderRepo is nil")
	}
	if orderDispatcher == nil {
		return nil, errors.New("orderDispatcher is nil")
	}

	return &AssignOrderHandler{orderRepo: orderRepo, courierRepo: courierRepo, unitOfWork: unitOfWork, orderDispatcher: orderDispatcher}, nil
}

func (ch *AssignOrderHandler) Handle(ctx context.Context, command AssignOrderCommand) error {
	if command.isEmpty() {
		return errors.New("empty command")
	}

	orderAg, err := ch.orderRepo.GetFirstCreatedOrder(ctx)
	if err != nil {
		return err
	}

	couriers, err := ch.courierRepo.GetAllFreeCouriers(ctx)
	if err != nil {
		return err
	}

	courier, err := ch.orderDispatcher.Dispatch(orderAg, couriers)
	if err != nil {
		return err
	}

	ctx = ch.unitOfWork.Begin(ctx)
	defer ch.unitOfWork.Rollback(ctx)

	err = ch.orderRepo.Update(ctx, orderAg)
	if err != nil {
		return err
	}
	err = ch.courierRepo.Update(ctx, courier)
	if err != nil {
		return err
	}

	err = ch.unitOfWork.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
