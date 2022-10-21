package main

import (
	"os"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/cmd/server/handler"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/docs"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/internal/products"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//IMPLEMENTAR MIDDLEWARE PARA TOKEN

//@title GoWeb MELI Bootcamp API
//@version 1.0
//@description This API handle products in a file
func main() {
	//leemos el archivo de variables de entorno
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	db := store.NewStore(store.FileType, "./products.json")

	repository := products.NewRepository(db)
	service := products.NewService(repository)
	p := handler.NewProductHandler(service)

	r := gin.Default()

	//generamos el endpoint para consultar la documentacion
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNameAndPrice())
	pr.DELETE("/:id", p.Delete())
	
	err = r.Run()
	if err != nil {
		panic(err)
	}
}