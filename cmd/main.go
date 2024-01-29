package main

import (
	"app/internal/application"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// env
	err := godotenv.Load(".local")
	if err != nil {
		fmt.Println(err)
		return
	}
	// ...

	// application
	// - config
	cfg := &application.ConfigDefault{
		Database: mysql.Config{
			User:      "root",
			Passwd:    os.Getenv("SERVER_PASSWD"),
			Net:       "tcp",
			Addr:      "127.0.0.1:3306",
			DBName:    "supermarket_db",
			ParseTime: true,
		},
		Address: "127.0.0.1:8080",
	}
	app := application.NewDefault(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
