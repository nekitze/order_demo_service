package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"orders_service/internal/models"
	"os"
)

type PostgresDatabase struct {
	DB *pg.DB
}

func NewPostgresDatabase() *PostgresDatabase {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	err := createSchema(db)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("successfully connected to database")
	return &PostgresDatabase{DB: db}
}

func createSchema(db *pg.DB) error {
	op := orm.CreateTableOptions{IfNotExists: true}
	err := db.Model(&models.Order{}).CreateTable(&op)
	if err != nil {
		return err
	}
	return nil
}
