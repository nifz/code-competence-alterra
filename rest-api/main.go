package main

import (
	"code-competence-alterra/rest-api/configs"
	"code-competence-alterra/rest-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = configs.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	routes.Init(e, db)

	e.Start(":8081")
}
