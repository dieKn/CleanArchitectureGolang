package controllers

import (
	"cleanArchitectureGolang/src/interfaceadaptors/presenters"
	usecases "cleanArchitectureGolang/src/usecases"
)

type HelloController struct {
	HelloWorldInteractor *usecases.HelloWorldInteractor
}

func NewHelloController(interactor *usecases.HelloWorldInteractor) *HelloController {
	return &HelloController{
		HelloWorldInteractor: interactor,
	}
}

func (controller *HelloController) GetHelloWorld() presenters.HelloWorldResponse {
	return presenters.ToResponse(controller.HelloWorldInteractor.GetHelloWorld("ハロー"))
}
