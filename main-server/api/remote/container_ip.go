package remote

import (
	"context"
	"github.com/aimamit/hostand/main-server/api/cache"
	"github.com/aimamit/hostand/main-server/proto"
)

func (p ProxyManger) GetContainerId(ctx context.Context, sd *proto.SubDomain) (*proto.IPResponseP, error) {
	s := cache.Cache.Get(sd.SubDomain)
	appVersion := &proto.AppVersion{
		Domain: sd.SubDomain,
		Version: s.String(),
	}
	res, err := InterServerCall.Docker.GetIPVersion(ctx, appVersion)
	if err != nil {
		return nil, err
	}

	return &proto.IPResponseP{Ip: res.Ip}, nil
}