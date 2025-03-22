package courierrepo

import (
	"lisichkinuriy/delivery/internal/adapters/out/postgres/shared"
	"lisichkinuriy/delivery/internal/domain/courier"
)

func DomainToDTO(aggregate *courier.Courier) CourierDTO {
	var courierDTO CourierDTO
	courierDTO.ID = aggregate.ID()
	courierDTO.Name = aggregate.Name().Value()
	courierDTO.Transport = TransportDTO{
		ID:        aggregate.Transport().ID(),
		Name:      aggregate.Transport().Name().Value(),
		Speed:     aggregate.Transport().Speed().Value(),
		CourierID: aggregate.ID(),
	}
	courierDTO.Location = shared.LocationDTO{
		X: aggregate.Location().X(),
		Y: aggregate.Location().Y(),
	}
	courierDTO.Status = aggregate.Status()
	return courierDTO
}

func DtoToDomain(dto CourierDTO) *courier.Courier {
	return courier.RestoreCourier(dto.ID, dto.Name, dto.Location.X, dto.Location.Y, dto.Status,
		dto.Transport.ID, dto.Transport.Name, dto.Transport.Speed)
}
