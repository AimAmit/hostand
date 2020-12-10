package remote

import (
	"context"
	"errors"
	"github.com/aimamit/hostand/main-server/proto"
	"log"
)

func (c ClientCall) ContainerCreate(domain, version string) error {
	appVersion := &proto.AppVersion{
		Domain: domain,
		Version: version,
	}

	res, err := c.Docker.ContainerCreate(context.Background(), appVersion)
	if err != nil {
		log.Println("Error while creating container: ", err)
		return err
	}
	if res.Error != "" {
		log.Println("Error while creating container: ", res.Error)
		return errors.New(res.Error)
	}

	return nil
}
