package drivers

import (
	"database/sql"
	"errors"
	"fmt"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MySQLConfig types.IMySQL

type mysqlProvider struct {
	debug bool
	DSN   string
	Gorm  *gorm.DB
	Sql   *sql.DB
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

	// Inject Gorm into repository
	for _, repository := range MySQLConfig.GetRepositories() {
		repository.SetDriver(mySQL.Gorm)
	}

	// Migration
	if MySQLConfig.Migration() {
		err := mySQL.Gorm.AutoMigrate(&migration{})
		if err != nil {
			lets.LogE("Unable to run migration %w", err)
			return
		}
		Migrate(mySQL.Gorm, mySQL.Sql)
	}
}

type migration struct {
	ID        uint   `gorm:"primaryKey,column:id"`
	Migration string `gorm:"column:migration"`
	Batch     uint   `gorm:"column:batch"`
}

func Migrate(g *gorm.DB, db *sql.DB) {
	// Define batch number
	var batch uint = 1
	lastMigration := &migration{}
	result := g.Last(lastMigration)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		lets.LogE("Unable to run migration %w", result.Error)
		return
	}

	batch = lastMigration.Batch + 1

	// Get migration files
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		// Search migration
		search := &migration{
			Migration: name,
		}

		result := g.Where("migration = ?", name).First(search)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			lets.LogE("Unable to run migration %w", result.Error)
			return
		}

		// Execute
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			lets.LogI("Migrating: %s", name)

			// Read file content
			filePath := fmt.Sprintf("migrations/%s", file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				lets.LogE("Unable to run migration: %s", err.Error())
				return
			}

			err = g.Transaction(func(tx *gorm.DB) error {
				for _, query := range strings.Split(string(content), ";") {
					query := strings.TrimSpace(query)
					if query == "" {
						continue
					}

					result = g.Exec(query)
					if result.Error != nil {
						return result.Error
					}
				}

				return nil
			})

			if err != nil {
				lets.LogE("Unable to run migration %w", err.Error())
				return
			}

			// Insert migration record
			m := &migration{
				Migration: name,
				Batch:     batch,
			}

			result = g.Create(m)
			if result.Error != nil {
				lets.LogE("Unable to run migration: %s", result.Error.Error())
				return
			}
		}

	}
}
