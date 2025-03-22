package orderrepo

import (
	"lisichkinuriy/delivery/internal/adapters/out/postgres/shared"
	"lisichkinuriy/delivery/internal/domain/order"
	"lisichkinuriy/delivery/internal/domain/vo"
)

func DomainToDTO(aggregate *order.Order) OrderDTO {
	var orderDTO OrderDTO
	orderDTO.ID = aggregate.ID()
	orderDTO.CourierID = aggregate.CourierID()
	orderDTO.Location = shared.LocationDTO{
		X: aggregate.Location().X(),
		Y: aggregate.Location().Y(),
	}
	orderDTO.Status = aggregate.Status()
	return orderDTO
}

func DtoToDomain(dto OrderDTO) *order.Order {
	var aggregate *order.Order
	location, _ := vo.NewLocation(dto.Location.X, dto.Location.Y)
	aggregate = order.RestoreOrder(dto.ID, dto.CourierID, location, dto.Status)
	return aggregate
}
