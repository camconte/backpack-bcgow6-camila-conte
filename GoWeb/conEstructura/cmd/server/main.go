package main

import (
	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/cmd/server/handler"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)
	p := handler.NewProductHandler(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	
	
	r.Run()
}