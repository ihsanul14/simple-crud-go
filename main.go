package main

import (
	"os"
	"simple-crud-go/database"
	"simple-crud-go/router"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	database.ConnectMySQL()
	router := router.InitRouter()
	router.Run(":" + os.Getenv("PORT"))
}
