package config

import "os"
import "fmt"
import "gorm.io/gorm"
import "gorm.io/driver/mysql"
import "github.com/joho/godotenv"

func SetupDatabaseConnection() *gorm.DB  {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Env file error")
	}
	
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB connection Error")
	}
	
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	
	if err != nil {
		panic("DB connection Error")
	}

	dbSQL.Close()
}