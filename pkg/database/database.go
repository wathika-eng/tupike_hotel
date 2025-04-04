package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"tupike_hotel/pkg/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var envs = config.Envs
var instance *service
var once sync.Once

type DBService interface {
	Health() map[string]string
	Close() error
	GetDB() *bun.DB
}

type service struct {
	db *bun.DB
}

func NewDatabase() (DBService, error) {
	var err error
	once.Do(func() {
		instance, err = initDB()
	})
	return instance, err
}

func initDB() (*service, error) {
	var dsn string
	if envs.ConnectionString == "" {
		dsn = fmt.Sprintf("%v://%v:%v@%v:%v/%v",
			envs.DbType, envs.DbUser, envs.DbPassword,
			envs.DbHost, envs.DbPort, envs.DbName)
	} else {
		dsn = envs.ConnectionString
	}
	log.Println(dsn)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	if sqldb == nil {
		return nil, fmt.Errorf("failed to create SQL DB connection")

	}
	if err := sqldb.Ping(); err != nil {
		sqldb.Close()
		return nil, fmt.Errorf("error connecting to the database: %v", err.Error())
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	log.Println("✅ Database connected successfully")
	return &service{
		db: db,
	}, nil
}

func (s *service) GetDB() *bun.DB {
	return s.db
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", s.db)
	return s.db.Close()
}
