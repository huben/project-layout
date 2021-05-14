package main

import (

	"layout/api/app/routers"
	"layout/api/app/routers/shop"
	"layout/api/app/routers/user"
	"layout/pkg/logger"
)

func main() {

	logger.InitLogger()

	defer logger.Sync()

	logger.Info("test = %s", 0)

	routers.Include(shop.Routers, user.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		logger.Info("err= %s", err)
	}
}
