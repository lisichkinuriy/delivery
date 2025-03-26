package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	ports "lisichkinuriy/delivery/internal/adapters/ports/mocks"
	"lisichkinuriy/delivery/internal/domain/order"
	"lisichkinuriy/delivery/internal/domain/vo"
	"testing"
)

func Test_CreateOrderHandler_Handle(t *testing.T) {
	ctx := context.Background()
	location := vo.MinLocation()
	orderAggregate, err := order.NewOrder(uuid.New(), location)
	require.NotNil(t, orderAggregate)

	orderRepoMock := &ports.IOrderRepositoryMock{}

	orderRepoMock.
		On("Get", ctx, orderAggregate.ID()).
		Return(nil, nil).
		Once()

	var captureObj *order.Order
	orderRepoMock.
		On("Add", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			captureObj = args.Get(1).(*order.Order)
		}).
		Return(nil).
		Once()

	handler, err := NewCreateOrderHandler(orderRepoMock)
	require.NoError(t, err)
	command, err := NewCreateOrderCommand(orderAggregate.ID(), "nostreet")
	require.NoError(t, err)

	err = handler.Handle(ctx, command)
	require.NoError(t, err)

	require.NotNil(t, orderRepoMock)
	require.NotNil(t, captureObj.ID())
	require.False(t, captureObj.Location().IsEmpty())
	require.True(t, captureObj.IsCreated())

}
