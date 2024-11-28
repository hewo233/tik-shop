package db

//import (
//	"github.com/joho/godotenv"
//	"gorm.io/driver/postgres"
//	"gorm.io/gen"
//	"gorm.io/gorm"
//	"log"
//)
//
//func main() {
//
//	env, err := godotenv.Read(".env")
//	if err != nil {
//		log.Fatal("Error reading .env file")
//	}
//
//	// 获取 DSN
//	dsn, exists := env["DATABASE_DSN"]
//	if !exists {
//		log.Fatal("DATABASE_DSN is not set in .env file")
//	}
//
//	// 初始化 GORM 数据库对象
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//	g := gen.NewGenerator(gen.Config{
//		OutPath: "./query",
//		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
//	})
//
//	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
//	g.UseDB(db)
//	g.ApplyBasic(
//		// Generate structs from all tables of current database
//		g.GenerateAllTable()...,
//	)
//	// Generate the code
//	g.Execute()
//}
