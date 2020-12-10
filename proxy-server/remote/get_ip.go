package remote

import (
	"context"
	"github.com/aimamit/hostand/proxy-server/proto"
)

func (c ServerCall)GetIP(subdomain string) (string, error){
	appVersion := &proto.AppVersionP{
		Domain: subdomain,
	}
	res, err := c.Main.GetContainerIp(context.Background(), appVersion)
	if err!= nil {
		return "", err
	}
	return res.Ip, nil
}
