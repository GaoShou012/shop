package microService

import "github.com/micro/go-micro/v2"

func Test() {
	service := micro.NewService(micro.Name("helloword!"))

	service.Init()
	service.Run()
}
