package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {

	migration := &migrate.FileMigrationSource{
		Dir: dir,
	}
	_, err := migrate.Exec(db.DB, "postgres", migration, migrate.Up)
	if err != nil {
		fmt.Println("Error running migrations:", err)
		return err
	}
	fmt.Println("Migrations applied successfully")
	return nil
}
