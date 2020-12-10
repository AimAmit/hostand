package api

import (
	"github.com/aimamit/hostand/main-server/api/cache"
	"github.com/aimamit/hostand/main-server/api/controller"
	"github.com/aimamit/hostand/main-server/api/remote"
)


func Run() {
	cache.GetRedisConn()
	remote.Run()
	controller.Run(":8000")
}
