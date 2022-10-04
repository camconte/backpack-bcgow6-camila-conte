package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type product struct{
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Colour string `json:"colour" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Stock int `json:"stock" binding:"required"`
	Code string `json:"code" binding:"required"`
	Published bool `json:"published" binding:"required"`
	CreatedAt string `json:"createdAt"`
}



//array en memoria
var products []product


func GetProducts() (products []product, err error){
	data, err0 := os.ReadFile("./products.json")
	if err0 != nil{
		err = err0
	}

	err1 := json.Unmarshal(data, &products)
	if err1 != nil{
		err = err1
	}

	return
}

func GetAll(c *gin.Context) {

	var products []product
	var err error

	products, err = GetProducts()

	if err != nil{
		c.JSON(500, err)
		return
	}

	var productsResponse []product

	//logica de filtrado
	productName := c.Query("name")
	productColour := c.Query("colour")
	productPrice := c.Query("price")
	productStock := c.Query("stock")
	productCreatedAt := c.Query("createdAt")

	if productName == "" && productColour == "" && productPrice == "" && productStock == "" && productCreatedAt == ""{
		c.JSON(200, products)
	}else{

		priceKey, _ := strconv.ParseFloat(productPrice, 64)
		stockKey, _ := strconv.ParseInt(productStock, 0, 0)


		for _, product  := range products {
			if (productName != "" && product.Name == productName) || (productColour != "" && product.Colour == productColour) || (productPrice != "" && product.Price == priceKey) || (productStock != "" && product.Stock == int(stockKey)) || (productCreatedAt != "" && product.CreatedAt == productCreatedAt)   {
				productsResponse = append(productsResponse, product)
			}
		}
	
		c.JSON(200, productsResponse)
	
	}

}

func GetById(c *gin.Context){
	products, err := GetProducts()
	if err != nil {
		c.JSON(500, err)
		return
	}

	var productResponse product

	idKey, err1 := strconv.ParseInt(c.Param("id"), 0, 0)
	if err1 != nil {
		c.JSON(500, err1)
		return
	}

	for _, product := range products {
		if product.Id == int(idKey){
			productResponse = product
		}
	}

	//se puede comparar con una estructura vacia
	productNull := product{
		Id: 0,
		Name: "",
		Colour: "",
		Price: 0,
		Stock: 0,
		Code: "",
		Published: false,
		CreatedAt: "",
	}

	if productResponse == productNull {
		c.String(404, "Product not found")
	}else{
		c.JSON(200, productResponse)
	}

}

func SaveProduct() gin.HandlerFunc{
	return func(c *gin.Context){
		var bodyReq product

		token := c.GetHeader("token")

		if token != "123456clc" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "you don't have permissions to make that request",
			})
			return
		}
		
		if err := c.ShouldBindJSON(&bodyReq); err != nil{
			
			//recorremos los errores detallados y los almacenamos para luego mostrarlos todos juntos
			errorMessages := []string{}

			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("the field %s is required", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}

		//tomo el id del producto en la ultima posicion dle array y lo incremento en 1 para asignarselo al nuevo producto
		bodyReq.Id = len(products) + 1
		//se formatea con esos numeros debido a las constantes con las que trabaja Go:
		//https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format

		bodyReq.CreatedAt = time.Now().Format("02-01-2006")

		products = append(products, bodyReq)

		c.JSON(http.StatusOK, bodyReq)

	}
}


func main(){
	//creamos el router
	router := gin.Default()

	//creamos el handler
	router.GET("/welcome", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hi Camila!",
		})
	})

	pr := router.Group("/products")

	//creamos la ruta para devolver los productos
	pr.GET("/", GetAll)

	//ruta para buscar por id
	pr.GET("/:id", GetById)

	//creamos el endpoint para almacenar un nuevo producto
	pr.POST("/save", SaveProduct())

	//corremos el servidor
	router.Run()
}