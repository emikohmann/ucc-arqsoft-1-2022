package controllers

import (
	"github.com/emikohmann/dependencies-injection/services"
	"testing"
)

func TestControllerImplementation_Util(t *testing.T) {

	var serviceMock services.ServiceMock

	var controllerImplementation ControllerImplementation
	controllerImplementation.Service = serviceMock

	controllerImplementation.Util()
}
