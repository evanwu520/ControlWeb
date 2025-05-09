package main

import (
	"game.com/controlWeb/db"
	"game.com/controlWeb/redis"
	"game.com/controlWeb/routers"
)

func main() {

	// init redis
	redis.InitRedis()
	// init db
	db.InitSQLServer()

	routers.SetupRouter().Run(":8080")

	//TODO graceful shutdown
}
