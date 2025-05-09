package db

import (
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var conn *gorm.DB

func NewSql() *newSql {
	return &newSql{}
}

type newSql struct {
}

func InitSQLServer() {

	dsn := "sqlserver://appuser:v%40hx1cZUn3Eu%2A%2B0pytWF@172.234.94.64:31433?database=PortalDB"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to SQL Server: %v", err)
		panic("failed to connect database")
	}
	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.DB: %v", err)
		panic("failed connection pool settings")
	}

	sqlDB.SetMaxOpenConns(5)                   // Maximum number of open connections
	sqlDB.SetMaxIdleConns(2)                   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Reuse connections for 30 minutes

	conn = db
}

func (*newSql) GetConn() *gorm.DB {

	return conn
}
