package http

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"lisichkinuriy/delivery/internal/adapters/in/http/problems"
	"lisichkinuriy/delivery/internal/application/usecases/commands"
	"lisichkinuriy/delivery/internal/application/usecases/queries"
	"lisichkinuriy/delivery/pkg/servers"
	"net/http"
)

type Server struct {
	CreateOrderHandler                commands.ICreateOrderHandler
	GetAllCouriersQueryHandler        queries.IGetAllCouriersQueryHandler
	GetNotCompletedOrdersQueryHandler queries.IGetNotCompletedOrdersQueryHandler
}

var _ servers.ServerInterface = &Server{}

func NewServer(
	CreateOrderHandler commands.ICreateOrderHandler,
	GetAllCouriersQueryHandler queries.IGetAllCouriersQueryHandler,
	GetNotCompletedOrdersQueryHandler queries.IGetNotCompletedOrdersQueryHandler,
) (*Server, error) {
	if CreateOrderHandler == nil {
		return nil, errors.New("CreateOrderHandler is nil")
	}
	if GetAllCouriersQueryHandler == nil {
		return nil, errors.New("GetAllCouriersQueryHandler is nil")
	}
	if GetNotCompletedOrdersQueryHandler == nil {
		return nil, errors.New("GetNotCompletedOrdersQueryHandler is nil")
	}

	return &Server{
		CreateOrderHandler:                CreateOrderHandler,
		GetAllCouriersQueryHandler:        GetAllCouriersQueryHandler,
		GetNotCompletedOrdersQueryHandler: GetNotCompletedOrdersQueryHandler,
	}, nil
}

func (s *Server) GetCouriers(ctx echo.Context) error {
	query, err := queries.NewGetAllCouriersQuery()
	if err != nil {
		return err
	}

	response, err := s.GetAllCouriersQueryHandler.Handle(query)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, problems.NewNotFound(err.Error()))
	}

	var couriers []servers.Courier
	for _, courier := range response.Couriers {
		location := servers.Location{
			X: courier.Location.X,
			Y: courier.Location.Y,
		}

		var courier = servers.Courier{
			Id:       courier.ID,
			Name:     courier.Name,
			Location: location,
		}
		couriers = append(couriers, courier)
	}
	return ctx.JSON(http.StatusOK, couriers)

}

func (s *Server) CreateOrder(ctx echo.Context) error {
	command, err := commands.NewCreateOrderCommand(uuid.New(), "None")
	if err != nil {
		return err
	}

	err = s.CreateOrderHandler.Handle(ctx.Request().Context(), command)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) GetOrders(ctx echo.Context) error {
	query, err := queries.NewGetNotCompletedOrdersQuery()
	if err != nil {
		return problems.NewBadRequest(err.Error())
	}

	response, err := s.GetNotCompletedOrdersQueryHandler.Handle(query)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, problems.NewNotFound(err.Error()))
	}

	var orders []servers.Order
	for _, courier := range response.Orders {
		location := servers.Location{
			X: courier.Location.X,
			Y: courier.Location.Y,
		}

		var courier = servers.Order{
			Id:       courier.ID,
			Location: location,
		}
		orders = append(orders, courier)
	}
	return ctx.JSON(http.StatusOK, orders)
}
