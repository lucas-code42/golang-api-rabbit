package db

import (
	"database/sql"
	"os"

	// driver mysql
	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB le as variaveis de ambiente e estebelece uma conexao com banco de dados
func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dataBase := os.Getenv("DB")
	urlConnection := user + ":" + pass + "@/" + dataBase + "?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
