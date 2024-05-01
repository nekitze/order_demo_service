package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"orders_service/internal/models"
)

type Database struct {
	DB *pg.DB
}

func NewDatabase() *Database {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "1234",
		Database: "postgres",
	})

	err := createSchema(db)
	if err != nil {
		log.Println(err)
	}

	return &Database{DB: db}
}

func createSchema(db *pg.DB) error {
	op := orm.CreateTableOptions{IfNotExists: true}
	err := db.Model(&models.Order{}).CreateTable(&op)
	if err != nil {
		return err
	}

	return nil
}
