package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

type Storage struct {
	*sql.DB
}

func New() *Storage {
	const op = "mysql.New"
	path := mustLoadEnv()

	db, err := sql.Open("mysql", path)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	return &Storage{db}
}

func mustLoadEnv() string {
	user, pass, host, dbName := loadEnv()
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbName)
}

func loadEnv() (user, pass, dbHost, dbName string) {
	const op = "mysql.loadEnv"

	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")

	return
}
