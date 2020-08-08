package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToDB(dbUser string, dbPasssword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPasssword, dbName,
	)
	// return gorm.Open("postgres", "host=ec2-54-204-37-92.compute-1.amazonaws.com port=5432 user=wfsobbobgfkelc dbname=d6fohcvmgj79rh password=785e47377465d41e2b515cd71cdbaf7ba72107b175d0d51c90f72bd2d22f14be")
}
