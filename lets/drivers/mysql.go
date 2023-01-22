package drivers

import (
	"database/sql"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MySQLConfig types.IMySQL

type mysqlProvider struct {
	debug   bool
	DSN     string
	Gorm    *gorm.DB
	Sql     *sql.DB
	Adapter func(*gorm.DB)
}

func (m *mysqlProvider) Connect() {
	var logType logger.Interface = logger.Default.LogMode(logger.Warn)
	if m.debug {
		logType = logger.Default.LogMode(logger.Info)
	}

	var err error
	m.Gorm, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.DSN, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL configs
	}), &gorm.Config{
		Logger: logType,
	})

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	m.Sql, err = m.Gorm.DB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	m.Sql.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	m.Sql.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	m.Sql.SetConnMaxLifetime(time.Hour)

}

// Define MySQL service host and port
func MySQL() {
	if MySQLConfig == nil {
		return
	}

	lets.LogI("MySQL Starting ...")

	mySQL := mysqlProvider{
		DSN:   MySQLConfig.GetDsn(),
		debug: MySQLConfig.DebugMode(),
	}
	mySQL.Connect()

	MySQLConfig.GetRepository().SetDriver(mySQL.Gorm)
}
