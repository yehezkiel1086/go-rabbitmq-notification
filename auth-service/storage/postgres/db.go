package postgres

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func InitPostgres(ctx context.Context, config *config.DB) (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Host, config.User, config.Password, config.DBName, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &DB{}, err
	}

	return &DB{
		db: db,
	}, nil
}

func (p *DB) GetDB() *gorm.DB {
	return p.db
}

func (p *DB) Migrate(dbs ...any) error {
	return p.db.AutoMigrate(dbs...)
}
