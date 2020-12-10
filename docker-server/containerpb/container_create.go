package containerpb

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"log"
	"os"
)

func (c *DockerClient)ContainerCreate(imageName, containerName string) error {

	ctx := context.Background()
	resp, err := c.Client.ContainerCreate(ctx,
		&container.Config{
			Image:        imageName,
			ExposedPorts: nat.PortSet{"3000": struct{}{},
			},
		}, &container.HostConfig{
			Resources: container.Resources{
				Memory: 32e+6,
				MemorySwap: 64e+6,
				NanoCPUs: 5e+7,
			},
			PortBindings: map[nat.Port][]nat.PortBinding{
				nat.Port("3000"): {{HostIP: "0.0.0.0", HostPort: "80"}},
			},
		},
		nil,
		nil,
		containerName,
	)
	if err != nil {
		return err
	}

	if err := c.Client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	out, err := c.Client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	w, err := stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	if err != nil {
		log.Println(w, err)
		return err
	}
	return nil
}
