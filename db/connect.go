package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	// 读取 .env 文件中的环境变量
	env, err := godotenv.Read("db/.env")
	if err != nil {
		log.Fatal("Error reading .env file")

	}

	// 获取数据库连接字符串 DSN
	dsn, exists := env["DATABASE_DSN"]
	if !exists {
		log.Fatal("DATABASE_DSN is not set in .env file")
		return nil, err
	}

	// 使用 GORM 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return db, nil
}
