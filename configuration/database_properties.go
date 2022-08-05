package configuration

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseProperties struct {
	username          string
	password          string
	host              string
	port              int
	databaseName      string
	driverName        string
	maxOpenConnection int
	maxIdleConnection int
	maxLifeTime       int
}

func NewDatabaseProperties() *DatabaseProperties {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse port with error %+v\n", err))
	}
	return &DatabaseProperties{
		username:     os.Getenv("DB_USERNAME"),
		password:     os.Getenv("DB_PASSWORD"),
		host:         os.Getenv("DB_HOST"),
		port:         port,
		databaseName: os.Getenv("DB_NAME"),
		driverName:   os.Getenv("DB_DRIVER"),
	}
}
