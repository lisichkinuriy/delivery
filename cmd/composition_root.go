package cmd

import (
	"lisichkinuriy/delivery/internal/domain/services"
)

type CompositionRoot struct {
	DomainServices DomainServices
}

type DomainServices struct {
	OrderDispatcher services.IOrderDispatcher
}

func NewCompositionRoot() CompositionRoot {

	orderDispatcher := services.NewOrderDispatcher()

	compositionRoot := CompositionRoot{
		DomainServices: DomainServices{
			OrderDispatcher: orderDispatcher,
		},
	}

	return compositionRoot
}
