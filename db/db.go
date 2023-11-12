package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DB_Connect() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "basico.db")

	if err != nil {
		log.Fatal("Erro na conexão com o basico!", err)
		return nil, err
	}
	return conn, nil
}
func DB_ConnectV2() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal("Erro na conexão com o bom!", err)
		return nil, err
	}
	return conn, nil
}
