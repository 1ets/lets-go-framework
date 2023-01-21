package types

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

const (
	MYSQL_DB_HOST       = "localhost"
	MYSQL_DB_PORT       = "3306"
	MYSQL_DB_USERNAME   = "root"
	MYSQL_DB_PASSWORD   = ""
	MYSQL_DB_DATABASE   = "lets"
	MYSQL_DB_CHARSET    = "utf8"
	MYSQL_DB_PARSE_TIME = "True"
	MYSQL_DB_LOC        = "Local"
)

type IMySQL interface {
	GetHost() string
	GetPort() string
	GetUsername() string
	GetPassword() string
	GetDatabase() string
	GetCharset() string
	GetParseTime() string
	GetLoc() string
	DebugMode() bool
	GetRepository() IMySQLRepository
	GetDsn() string
}

type MySQL struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Charset    string
	ParseTime  string
	Loc        string
	Debug      bool
	Gorm       *gorm.DB
	DB         *sql.DB
	Repository IMySQLRepository
}

func (mysql *MySQL) GetHost() string {
	if mysql.Host == "" {
		fmt.Println("Configs MySQL: DB_HOST is not set in .env file, using default configuration.")
		return MYSQL_DB_HOST
	}
	return mysql.Host
}

func (mysql *MySQL) GetPort() string {
	if mysql.Host == "" {
		fmt.Println("Configs MySQL: DB_PORT is not set in .env file, using default configuration.")
		return MYSQL_DB_PORT
	}
	return mysql.Port
}

func (mysql *MySQL) GetUsername() string {
	if mysql.Username == "" {
		fmt.Println("Configs MySQL: DB_USERNAME is not set in .env file, using default configuration.")
		return MYSQL_DB_USERNAME
	}
	return mysql.Username
}

func (mysql *MySQL) GetPassword() string {
	if mysql.Host == "" {
		fmt.Println("Configs MySQL: DB_PASSWORD is not set in .env file, using default configuration.")
		return MYSQL_DB_PASSWORD
	}
	return mysql.Password
}

func (mysql *MySQL) GetDatabase() string {
	if mysql.Database == "" {
		fmt.Println("Configs MySQL: DB_DATABASE is not set in .env file, using default configuration.")
		return MYSQL_DB_DATABASE
	}
	return mysql.Database
}

func (mysql *MySQL) GetCharset() string {
	if mysql.Charset == "" {
		fmt.Println("Configs MySQL: Charset is not set in configs, using default configuration.")
		return MYSQL_DB_CHARSET
	}
	return mysql.Charset
}

func (mysql *MySQL) GetParseTime() string {
	if mysql.ParseTime == "" {
		fmt.Println("Configs MySQL: Charset is not set in configs, using default configuration.")
		return MYSQL_DB_CHARSET
	}
	return mysql.ParseTime
}

func (mysql *MySQL) GetLoc() string {
	if mysql.Loc == "" {
		fmt.Println("Configs MySQL: Loc is not set in configs, using default configuration.")
		return MYSQL_DB_LOC
	}
	return mysql.Loc
}

func (mysql *MySQL) DebugMode() bool {
	return mysql.Debug
}

func (mysql *MySQL) GetRepository() IMySQLRepository {
	return mysql.Repository
}

func (mysql *MySQL) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		mysql.GetUsername(),
		mysql.GetPassword(),
		mysql.GetHost(),
		mysql.GetPort(),
		mysql.GetDatabase(),
		mysql.GetCharset(),
		mysql.GetParseTime(),
		mysql.GetLoc(),
	)
}
