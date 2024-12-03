package main

import (
	"github.com/hewo/tik-shop/db/model"
	"gorm.io/gen"
)

func main() {
	//
	//// 初始化 GORM 数据库对象
	//db, err := connectDB.ConnectDB()
	//if err != nil {
	//	panic(err)
	//}
	//err = db.AutoMigrate(
	//	&model.Users{},
	//	&model.Product{},
	//	&model.Order{},
	//	&model.Address{},
	//	&model.CartItem{},
	//	&model.OrderItem{},
	//	&model.PaymentDetails{})
	//if err != nil {
	//	panic(err)
	//}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	//g.UseDB(db)
	g.ApplyBasic(
		&model.Users{},
		&model.Product{},
		&model.Order{},
		&model.Address{},
		&model.CartItem{},
		&model.OrderItem{},
		&model.PaymentDetails{},
	)
	// Generate the code
	g.Execute()
}
