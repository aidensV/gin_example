package main

import (
	"log"
	// "os"

	"github.com/aidensV/gin_example/configs"
	"github.com/aidensV/gin_example/database"
	"github.com/aidensV/gin_example/models"
	"github.com/aidensV/gin_example/repositories"
)

func main() {
	//database config
	dbUser, dbPassword, dbName := "root", "@Crossgg52b", "db_go_inven"
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

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }
	route.Run(":" + "8080")
}
