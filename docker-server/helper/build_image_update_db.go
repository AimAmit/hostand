package helper

import (
	"bufio"
	"github.com/aimamit/hostand/docker-server/containerpb"
	"io"
)

func BuildImageAndUpdateDb(domain, version string, containerTarFile io.Reader){
	go containerpb.DockerCli.BuildImage(domain, version, bufio.NewReader(containerTarFile))
}
