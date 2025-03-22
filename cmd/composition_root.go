package cmd

import (
	"context"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/internal/adapters/out/postgres"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/courierrepo"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/orderrepo"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/services"
	"log"
)

type CompositionRoot struct {
	DomainServices DomainServices
	Repositories   Repositories
}

type DomainServices struct {
	OrderDispatcher services.IOrderDispatcher
}

type Repositories struct {
	UnitOfWork        ports.IUnitOfWork
	OrderRepository   ports.IOrderRepository
	CourierRepository ports.ICourierRepository
}

func NewCompositionRoot(ctx context.Context, db *gorm.DB) CompositionRoot {

	orderDispatcher := services.NewOrderDispatcher()

	unitOfWork, err := postgres.NewUnitOfWork(db)
	if err != nil {
		log.Fatal(err)
	}
	orderRepo, err := orderrepo.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	courierRepo, err := courierrepo.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	compositionRoot := CompositionRoot{
		DomainServices: DomainServices{
			OrderDispatcher: orderDispatcher,
		},
		Repositories: Repositories{
			UnitOfWork:        unitOfWork,
			OrderRepository:   orderRepo,
			CourierRepository: courierRepo,
		},
	}

	return compositionRoot
}
