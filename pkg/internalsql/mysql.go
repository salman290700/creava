package internalsql

import (
	"database/sql"
	"fmt"
	"gotweet/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySql(cfg *config.Config) (*sql.DB, error) {
	dataSourcename := fmt.Sprintf(("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s"), cfg.Db_user, cfg.Db_password, cfg.Db_host, cfg.Db_port, cfg.Db_name, "Asia%2FJakarta")
	db, err := sql.Open("mysql", dataSourcename)
	if err != nil {
		return nil, fmt.Errorf("error connection to database")
	}
	log.Println("database connected")
	return db, nil
}
