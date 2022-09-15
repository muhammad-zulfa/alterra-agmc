package main

import (
	"github.com/muhammad-zulfa/alterra-agmc/controllers"
	"github.com/muhammad-zulfa/alterra-agmc/db"
	"github.com/muhammad-zulfa/alterra-agmc/models"
	"github.com/muhammad-zulfa/alterra-agmc/routes"
)

func main() {
	db.Init()
	db := db.DbManager()

	db.AutoMigrate(models.Book{})

	bookController := controllers.NewBookController()

	e := routes.Init(bookController)

	e.Logger.Fatal(e.Start(":1323"))
}
