package db

import (
	"fmt"

	"github.com/Likhon22/ecom/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return config.GetConfig().DBUrl
}

func NewConnection() (*sqlx.DB, error) {
	dbsource := GetConnectionString()
	db, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		fmt.Println("failed to connect to db:", err)
		return nil, err
	}
	return db, nil

}
