package main

import (
	"github.com/hewo/tik-shop/db/connectDB"
	"github.com/hewo/tik-shop/db/model"
	"gorm.io/gen"
)

func main() {

	// 初始化 GORM 数据库对象
	db, err := connectDB.ConnectDB()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		&model.Address{},
		&model.CartItem{},
		&model.Order{},
		&model.OrderItem{},
		&model.PaymentDetails{},
		&model.Product{},
		&model.Users{})
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)
	g.ApplyBasic(
		&model.Address{},
		&model.CartItem{},
		&model.Order{},
		&model.OrderItem{},
		&model.PaymentDetails{},
		&model.Product{},
		&model.Users{},
	)
	// Generate the code
	g.Execute()
}
