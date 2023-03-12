package main

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/bloberror"
	"log"
)

var blobEndpoint = "http://127.0.0.1:10000/devstoreaccount1"
var connectionString = "DefaultEndpointsProtocol=http;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;BlobEndpoint=http://127.0.0.1:10000/devstoreaccount1;"
var containers = []string{"thumbnails"}

func NewAzBlobClient() *azblob.Client {
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func SeedContainers(ctx context.Context, client *azblob.Client) {
	for _, container := range containers {
		access := azblob.PublicAccessTypeContainer
		_, err := client.CreateContainer(ctx, container, &azblob.CreateContainerOptions{
			Access:       &access,
			Metadata:     nil,
			CPKScopeInfo: nil,
		})
		switch err {
		case nil:
			log.Printf("az blob storage container created: %s", container)
		default:
			if bloberror.HasCode(err, bloberror.ContainerAlreadyExists) {
				log.Printf("az blob storage container already exists: %s", container)
			} else {
				log.Fatal(err)
			}
		}
	}
}
