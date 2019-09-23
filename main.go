package main

import (
	"log"

	"github.com/aidensV/gin_example/configs"
	"github.com/aidensV/gin_example/database"
	"github.com/aidensV/gin_example/models"
	"github.com/aidensV/gin_example/repositories"
)

func main() {
	//database config
	dbUser, dbPassword, dbName := "postgres", "postgres", "gorm_crud_example"
	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	//unable to connect to database
	if err != nil {
		log.Fatalln(err)
	}

	//ping to database
	err = db.DB().Ping()

	//error ping to database
	if err != nil {
		log.Fatalln(err)
	}
	//migration
	db.AutoMigrate(&models.Contact{})

	defer db.Close()

	contactRepository := repositories.NewContactRepository(db)
	route := configs.SetupRoutes(contactRepository)

	route.Run(":8000")
}
