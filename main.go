
package main

import (
	"github.com/micro/go-micro"
	proto "github.com/artback/kafka_microservices_go/greeter"
	"golang.org/x/net/context"
	"log"
)
type Greeter struct{}
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)        service.Init()        proto.RegisterGreeterHandler(service.Server(), new(Greeter))        if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}