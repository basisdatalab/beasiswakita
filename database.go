package beasiswakita

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var DbMap *gorp.DbMap

var Transaction *gorp.Transaction

func InitDb(dsn string) error {
	db, _ := sql.Open("mysql", fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", dsn))
	err := db.Ping()
	if err != nil {
		return err
	}

	database := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	database.AddTableWithName(User{}, "users").SetKeys(true, "ID")
	database.AddTableWithName(Organization{}, "organizations").SetKeys(true, "ID")
	database.AddTableWithName(Student{}, "students").SetKeys(true, "ID")
	DbMap = database

	return nil
}
