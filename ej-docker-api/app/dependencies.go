package app

import (
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/config"
	itemsHandler "github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/controllers/items"
	itemsService "github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/services/items"
	"log"
)

type Dependencies interface {
	ItemsSvc() itemsService.Service
	ItemsHdl() itemsHandler.Handler
}

func GetDependencies() Dependencies {
	if config.IsTest() {
		log.Println("setting up mock dependencies")
		return NewDependenciesMock()
	}
	log.Println("setting up prod dependencies")
	return NewDependenciesProd()
}
