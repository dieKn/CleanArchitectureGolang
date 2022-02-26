package model

type HelloWorld struct {
	Hello string
	World *World
}

func (hw *HelloWorld) CorrectHelloWorld(world string) string {
	if hw.World != nil && hw.World.Value != nil {
		return hw.Hello + *hw.World.Value
	}
	return hw.Hello + world
}
