package db

import (
	"ForumPublica/server/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var (
	DB *xorm.Engine = nil
)

func Migrate() error {
	m, err := migrate.New("file://server/migrations", "mysql://"+config.Vars.DBC+"?multiStatements=true")
	if err != nil {
		return err
	}

	errSt := m.Up()
	if errSt != nil && errSt != migrate.ErrNoChange {
		return errSt
	}

	return nil

}

func Connect() {
	var err error
	DB, err = xorm.NewEngine("mysql", config.Vars.DBC)
	DB.ShowSQL(true)

	if err != nil {
		fmt.Println("Failed to connect database:", err)
	}
}
