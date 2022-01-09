package main

import (
	"context"
	"log"
	"net"

	"github.com/shamaton/grpc-msgpackgen-sample/encoding"
	"github.com/shamaton/grpc-msgpackgen-sample/pb"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedGreeterServer
}

func (s *service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("call from", req.Name)
	rsp := new(pb.HelloReply)
	rsp.Message = "Hello " + req.Name + "."
	return rsp, nil
}

func main() {

	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}

	encoding.RegisterCodec()

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &service{})
	s.Serve(l)

}
