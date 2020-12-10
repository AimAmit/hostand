package api

import (
	"context"
	"fmt"
	"github.com/aimamit/hostand/docker-server/containerpb"
	"github.com/aimamit/hostand/docker-server/proto"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (c GrpcManger) GetIPVersion(ctx context.Context, appVersion *proto.AppVersion) (*proto.IPResponse, error) {

	domain := appVersion.Domain
	version := appVersion.Version
	imageName := fmt.Sprintf("%s:%s", domain, version)

	args := filters.NewArgs()
	args.Add("ancestor", imageName)
	containers, err := containerpb.DockerCli.Client.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: args,
	})

	if err != nil {
		return &proto.IPResponse{Ip: ""}, err
	}

	containerJson, err := containerpb.DockerCli.Client.ContainerInspect(context.Background(), containers[0].ID)
	if err != nil {
		return &proto.IPResponse{Ip: ""}, err
	}

	return &proto.IPResponse{Ip: containerJson.NetworkSettings.IPAddress}, nil
}
