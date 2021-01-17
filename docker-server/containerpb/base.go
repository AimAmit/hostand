package containerpb

import "github.com/docker/docker/client"

type DockerClient struct {
	Client client.APIClient
}

var DockerCli DockerClient

func DockerClientInit() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	DockerCli.Client = cli
}