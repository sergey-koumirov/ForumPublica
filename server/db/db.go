package db

import (
	"ForumPublica/server/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var (
	DB *gorm.DB = nil
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
	DB, err = gorm.Open("mysql", config.Vars.DBC)
	DB.LogMode(true)

	if err != nil {
		fmt.Println("Failed to connect database:", err)
	}
}
