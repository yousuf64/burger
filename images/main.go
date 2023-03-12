package main

import (
	"context"
	"fmt"
	"github.com/yousuf64/burger/protos/images"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var addr = "localhost:3001"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	azBlobClient := NewAzBlobClient()
	SeedContainers(ctx, azBlobClient)
	imgSvr := imagesServer{azBlob: azBlobClient}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listening on %s\n", addr)

	var opts []grpc.ServerOption
	grpcSvr := grpc.NewServer(opts...)
	images.RegisterImagesServer(grpcSvr, imgSvr)
	err = grpcSvr.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
