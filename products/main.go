package main

import (
	"context"
	"fmt"
	"github.com/yousuf64/burger/protos/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type productsServer struct {
	products.UnimplementedProductsServer
}

func (productsServer) CreateProduct(context.Context, *products.CreateProductRequest) (*products.CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (productsServer) GetProduct(context.Context, *products.GetProductRequest) (*products.GetProductResponse, error) {
	product := &products.GetProductResponse{
		Name:        "big mac",
		Sku:         "SKU911",
		Description: "some massive burger",
		Price:       1299,
		Thumbnail:   "",
		Available:   true,
		Featured:    true,
		Images:      nil,
		CategoryId:  99,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		DeletedAt:   0,
	}
	return product, nil
}
func (productsServer) GetProducts(context.Context, *products.GetProductsRequest) (*products.GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}

func main() {
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on localhost:3000")

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	products.RegisterProductsServer(server, productsServer{})
	server.Serve(lis)
}
