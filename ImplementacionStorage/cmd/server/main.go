package main

import (
	"clase1/cmd/server/handler"
	"clase1/pkg/db"
	"clase1/internal/products"
)

func main() {


	engine, db := db.ConnectDatabase()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)


	pr := engine.Group("/api/v1/products")

	pr.POST("/", p.Store())
	//recibe el name mediante el body
	pr.GET("/", p.GetByName())

	pr.GET("/all", p.GetAll())

	pr.GET("/:id", p.GetProductsByWarehouse())

	pr.DELETE("/:id", p.Delete())

	engine.Run()
}


