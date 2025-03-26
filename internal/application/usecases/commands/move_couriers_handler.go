package commands

import (
	"context"
	"errors"
	"lisichkinuriy/delivery/internal/adapters/ports"
)

type IMoveCouriersHandler interface {
	Handle(context.Context, MoveCouriersCommand) error
}

var _ IMoveCouriersHandler = &MoveCouriersHandler{}

type MoveCouriersHandler struct {
	unitOfWork  ports.IUnitOfWork
	courierRepo ports.ICourierRepository
	orderRepo   ports.IOrderRepository
}

func NewMoveCouriersHandler(
	unitOfWork ports.IUnitOfWork,
	courierRepo ports.ICourierRepository,
	orderRepo ports.IOrderRepository,
) (*MoveCouriersHandler, error) {

	if unitOfWork == nil {
		return nil, errors.New("unitOfWork is nil")
	}
	if courierRepo == nil {
		return nil, errors.New("courierRepo is nil")
	}
	if orderRepo == nil {
		return nil, errors.New("orderRepo is nil")
	}

	return &MoveCouriersHandler{orderRepo: orderRepo, courierRepo: courierRepo, unitOfWork: unitOfWork}, nil
}

func (ch *MoveCouriersHandler) Handle(ctx context.Context, command MoveCouriersCommand) error {
	if command.isEmpty() {
		return errors.New("empty command")
	}

	assignedOrders, err := ch.orderRepo.GetAssignedOrders(ctx)
	if err != nil {
		return err
	}

	ctx = ch.unitOfWork.Begin(ctx)
	defer ch.unitOfWork.Rollback(ctx)

	for _, assignedOrder := range assignedOrders {
		courier, err := ch.courierRepo.Get(ctx, *assignedOrder.CourierID())
		if err != nil {
			return err
		}

		err = courier.Move(assignedOrder.Location())
		if err != nil {
			return err
		}

		if courier.Location().Equals(assignedOrder.Location()) {
			err := assignedOrder.Complete()
			if err != nil {
				return err
			}

			err = courier.SetFree()
			if err != nil {
				return err
			}
		}

		err = ch.orderRepo.Update(ctx, assignedOrder)
		if err != nil {
			return err
		}
		err = ch.courierRepo.Update(ctx, courier)
		if err != nil {
			return err
		}
	}

	err = ch.unitOfWork.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
