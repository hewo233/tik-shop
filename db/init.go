package main

import (
	"github.com/hewo/tik-shop/db/connectDB"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/shared/consts"
	"gorm.io/gen"
)

func main() {

	//初始化 GORM 数据库对象
	db, err := connectDB.ConnectDB(consts.InitDBEnvPath)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.CartItem{},
		&model.OrderItem{},
	)

	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)
	g.ApplyBasic(
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.CartItem{},
		&model.OrderItem{},
	)
	// Generate the code
	g.Execute()
}
