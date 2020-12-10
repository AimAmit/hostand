package api

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/aimamit/hostand/docker-server/containerpb"
	"github.com/aimamit/hostand/docker-server/proto"
	"github.com/google/uuid"
	"io"
	"log"
)

func (c GrpcManger) FileUpload(server proto.DockerService_FileUploadServer) error {

	containerTarFile := bytes.Buffer{}

	stream, _ := server.Recv()
	appVersion := stream.GetAppVersion()
	domain := appVersion.GetDomain()
	version := appVersion.GetVersion()

	for {
		stream, err := server.Recv()
		if err == io.EOF {
			log.Println("file received")
			containerpb.DockerCli.BuildImage(domain, version, bufio.NewReader(&containerTarFile))
			return server.SendAndClose(&proto.FileResponse{Error: ""})
		}
		if err != nil {
			return server.SendAndClose(&proto.FileResponse{Error: fmt.Sprintf("%v", err)})
		}

		chunk := stream.GetChunk()

		_, err = containerTarFile.Write(chunk)
		if err != nil {
			return server.SendAndClose(&proto.FileResponse{Error: ""})
		}
	}
}

func (c GrpcManger) ContainerCreate(ctx context.Context, appVersion *proto.AppVersion) (*proto.FileResponse, error) {
	domain := appVersion.Domain
	version := appVersion.Version
	imageName := fmt.Sprintf("%s:%s", domain, version)
	containerName := fmt.Sprintf("%s.%s.%s", domain, version, uuid.New())
	err := containerpb.DockerCli.ContainerCreate(imageName, containerName)
	if err != nil {
		return &proto.FileResponse{Error: err.Error()}, err
	}
	return &proto.FileResponse{Error: ""}, nil
}

