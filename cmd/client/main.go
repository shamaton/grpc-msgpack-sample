package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/shamaton/grpc-msgpackgen-sample/encoding"

	"github.com/shamaton/grpc-msgpackgen-sample/pb"
	"google.golang.org/grpc"
)

func main() {

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.CallContentSubtype(encoding.JsonCodecName)),
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:12345", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
