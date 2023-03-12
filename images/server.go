package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/yousuf64/burger/protos/images"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type imagesServer struct {
	images.UnimplementedImagesServer

	azBlob *azblob.Client
}

func (img imagesServer) UploadImage(ctx context.Context, in *images.UploadImageRequest) (*images.UploadImageResponse, error) {
	_, err := img.azBlob.UploadBuffer(ctx, in.ContainerName, in.FileName, in.Content, nil)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to upload image: %v", err)
	}

	return &images.UploadImageResponse{Url: fmt.Sprintf("%s/thumbnails/%s", blobEndpoint, in.FileName)}, nil
}
