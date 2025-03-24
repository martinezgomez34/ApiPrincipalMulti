package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error verificando la conexión a la base de datos:", err)
	}

	log.Println("Conexión a la base de datos establecida")

	return &Database{DB: db}
}

func (d *Database) Close() {
	d.DB.Close()
	log.Println("Conexión a la base de datos cerrada")
}