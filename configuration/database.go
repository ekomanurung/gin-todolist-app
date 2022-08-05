package configuration

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbEngine *Database

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	if dbEngine == nil {
		dbEngine = &Database{}
	}
	return dbEngine
}

func (d *Database) GetDBEngine() *Database {
	return dbEngine
}

func (d *Database) InitializeConnection(props *DatabaseProperties) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		props.username, props.password, props.host, props.port, props.databaseName)

	conn, err := sql.Open(props.driverName, dsn)
	if err != nil {
		panic(fmt.Sprintf("Panic when initialize mysql driver connection caused by: %+v\n", err))
	}

	conn.SetMaxOpenConns(props.maxOpenConnection)
	conn.SetMaxIdleConns(props.maxIdleConnection)
	conn.SetConnMaxLifetime(time.Duration(props.maxLifeTime) * time.Minute)

	//TODO switch case if driver name is not mysql then change configuration
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: conn}))
	if err != nil {
		panic(fmt.Sprintf("Panic when initialize Gorm db caused by: %+v\n", err))
	}

	d.db = db
}

func (d *Database) Migrate(v interface{}) {
	err := d.GetDBEngine().db.AutoMigrate(v)
	if err != nil {
		panic(fmt.Sprintf("Panic when run db migration caused by: %+v\n", err))
	}
}
