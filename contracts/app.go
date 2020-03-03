package contracts

import (
	"app/bootstrap"
)

type (
	App struct {
		bootstrap.MySQL
		bootstrap.RedisDB
	}
)
