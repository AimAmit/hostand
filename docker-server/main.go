package main

import (
	"github.com/aimamit/hostand/docker-server/api"
	"github.com/aimamit/hostand/docker-server/containerpb"
)

func main() {

	containerpb.DockerClientInit()
	api.GrpcServerInit()
}