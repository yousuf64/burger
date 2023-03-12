terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.1"
    }
  }
}

provider "docker" {}

# Pull image
resource "docker_image" "azurite" {
  name         = "mcr.microsoft.com/azure-storage/azurite:latest"
  keep_locally = true
}

# Create container
resource "docker_container" "azurite" {
  image = docker_image.azurite.image_id
  name  = "azure-blob"
  ports {
    internal = 10000
    external = 10000
  }
  command = ["azurite-blob", "--blobHost", "0.0.0.0", "--blobPort", "10000"]
}