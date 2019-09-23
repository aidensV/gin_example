package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToDB(dbUser string, dbPasssword string, dbName string) (*gorm.DB, error) {
	// var connectionString = fmt.Sprintf(
	// 	"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	dbUser, dbPasssword, dbName,
	// )
	return gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm_crud_example password=postgres")
}
