package cmd

import (
	"context"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/internal/adapters/in/jobs"
	"lisichkinuriy/delivery/internal/adapters/out/grpc"
	"lisichkinuriy/delivery/internal/adapters/out/postgres"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/courierrepo"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/orderrepo"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/application/usecases/commands"
	"lisichkinuriy/delivery/internal/application/usecases/queries"
	"lisichkinuriy/delivery/internal/domain/services"
	"log"
)

type CompositionRoot struct {
	DomainServices  DomainServices
	Repositories    Repositories
	CommandHandlers CommandHandlers
	QueryHandlers   QueryHandlers
	Jobs            Jobs
	Clients         Clients
}

type Clients struct {
	Geo ports.IGeoClient
}

type Jobs struct {
	MoveCouriersJob cron.Job
	AssignOrdersJob cron.Job
}

type CommandHandlers struct {
	AssignOrderHandler  commands.IAssignOrderHandler
	CreateOrderHandler  commands.ICreateOrderHandler
	MoveCouriersHandler commands.IMoveCouriersHandler
}

type QueryHandlers struct {
	GetAllCouriersQueryHandler        queries.IGetAllCouriersQueryHandler
	GetNotCompletedOrdersQueryHandler queries.IGetNotCompletedOrdersQueryHandler
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

	geoclient, err := grpc.NewGRPCGeoClient("localhost:5004")
	if err != nil {
		log.Fatalf("could not create grpc geoclient: %v", err)
	}

	// Command Handlers
	createOrderHandler, err := commands.NewCreateOrderHandler(orderRepo, geoclient)
	if err != nil {
		log.Fatalf("run application error: %s", err)
	}

	assignOrderHandler, err := commands.NewAssignOrderHandler(
		unitOfWork, courierRepo, orderRepo, orderDispatcher)
	if err != nil {
		log.Fatalf("run application error: %s", err)
	}

	moveCouriersHandler, err := commands.NewMoveCouriersHandler(
		unitOfWork, courierRepo, orderRepo)
	if err != nil {
		log.Fatalf("run application error: %s", err)
	}

	// Query Handlers
	getAllCouriersQueryHandler, err := queries.NewGetAllCouriersQueryHandler(db)
	if err != nil {
		log.Fatalf("run application error: %s", err)
	}

	getNotCompletedOrdersQueryHandler, err := queries.NewGetNotCompletedOrdersQueryHandler(db)
	if err != nil {
		log.Fatalf("run application error: %s", err)
	}

	assignOrderJob, err := jobs.NewAssignOrdersJob(assignOrderHandler)
	if err != nil {
		log.Fatalf("run application error: %s", err)
	}

	moveCouriersJob, err := jobs.NewMoveCouriersJob(moveCouriersHandler)
	if err != nil {
		log.Fatalf("run application error: %s", err)
	}

	compositionRoot := CompositionRoot{
		DomainServices: DomainServices{
			orderDispatcher,
		},
		Repositories: Repositories{
			unitOfWork,
			orderRepo,
			courierRepo,
		},
		CommandHandlers: CommandHandlers{
			assignOrderHandler,
			createOrderHandler,
			moveCouriersHandler,
		},
		QueryHandlers: QueryHandlers{
			getAllCouriersQueryHandler,
			getNotCompletedOrdersQueryHandler,
		},
		Jobs: Jobs{
			moveCouriersJob,
			assignOrderJob,
		},
		Clients: Clients{
			geoclient,
		},
	}

	return compositionRoot
}
