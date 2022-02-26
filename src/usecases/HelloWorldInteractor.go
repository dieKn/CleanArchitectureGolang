package usecases

import (
	"cleanArchitectureGolang/src/entities/model"
	"cleanArchitectureGolang/src/entities/repository"
)

type HelloWorldInteractor struct {
	WorldRepository *repository.WorldRepository
}

func NewHelloWorldInteractor(
// FIXME: GateWayに繋ぐ際に実装する。
//repository repository.WorldRepository
) *HelloWorldInteractor {
	return &HelloWorldInteractor{
		// WorldRepository: repository,
	}
}

func (interactor *HelloWorldInteractor) GetHelloWorld(hello string) string {
	// TODO: 現状DBと繋ぐまでは実装しないので一旦nilを返す。
	// world := interactor.WorldRepository.GetWorld()
	helloWorld := model.HelloWorld{
		hello,
		// world,
		nil,
	}
	return helloWorld.CorrectHelloWorld("ワールド！")

}
