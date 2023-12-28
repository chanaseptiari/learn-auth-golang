package initializer

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	switch os.Getenv("DRIVER") {
	case "mysql":
		DB = ConnectMysql(os.Getenv("DSN"))
	case "postgres":
		DB = ConnectPostgres(os.Getenv("DSN"))
	default:
		panic("please set mysql or postgres")
	}
	config, err := DB.DB()
	if err != nil {
		panic(err)
	}
	config.SetMaxOpenConns(100)
	config.SetMaxIdleConns(10)
	config.SetConnMaxLifetime(30 * 60)
	config.SetConnMaxIdleTime(5 * 60)

	return DB
}

func ConnectMysql(dsn string) *gorm.DB {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func ConnectPostgres(dsn string) *gorm.DB {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return conn
}
