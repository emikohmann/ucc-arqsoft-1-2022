package app

import (
	itemsHandler "github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/controllers/items"
	itemsService "github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/services/items"
)

type DependenciesProd struct {
	itemsSvc itemsService.Service
	itemsHdl itemsHandler.Handler
}

func NewDependenciesProd() DependenciesProd {
	itemsSvc := itemsService.NewServiceImpl()
	itemsHdl := itemsHandler.NewHandlerImpl(itemsSvc)

	return DependenciesProd{
		itemsSvc: itemsSvc,
		itemsHdl: itemsHdl,
	}
}

func (dep DependenciesProd) ItemsSvc() itemsService.Service {
	return dep.itemsSvc
}

func (dep DependenciesProd) ItemsHdl() itemsHandler.Handler {
	return dep.itemsHdl
}
