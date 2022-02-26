package presenters

type HelloWorldResponse struct {
	HelloWorld string
}

func ToResponse(entity string) HelloWorldResponse {
	return HelloWorldResponse{
		HelloWorld: entity,
	}
}
