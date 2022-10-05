package main

import (
	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/cmd/server/handler"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//leemos el archivo de variables de entorno
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}

	repository := products.NewRepository()
	service := products.NewService(repository)
	p := handler.NewProductHandler(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNameAndPrice())
	pr.DELETE("/:id", p.Delete())
	
	r.Run()
}