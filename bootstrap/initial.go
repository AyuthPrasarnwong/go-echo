package bootstrap

func InitialDatabases() {
	CreateMySQLConnection()
	CreateRedisConnection()
}
