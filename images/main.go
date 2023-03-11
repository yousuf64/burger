package main

import (
	"context"
	"fmt"
	"github.com/yousuf64/burger/protos/images"
	"google.golang.org/grpc"
	"log"
	"net"
)

type imagesServer struct {
	images.UnimplementedImagesServer
}

func (imagesServer) UploadImage(context.Context, *images.UploadImageRequest) (*images.UploadImageResponse, error) {
	return &images.UploadImageResponse{
		Url: "image-server/foo.jpg",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:3001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on localhost:3001")

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	images.RegisterImagesServer(server, imagesServer{})
	server.Serve(lis)
}
