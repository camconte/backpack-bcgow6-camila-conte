package main

import (
	"encoding/json"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

type product struct{
	Id int
	Name string
	Colour string
	Price float64
	Stock int
	Code string
	Published bool
	CreatedAt string
}

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


func main(){
	//creamos el router
	router := gin.Default()

	//creamos el handler
	router.GET("/welcome", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hi Camila!",
		})
	})

	//creamos la ruta para devolver los productos
	router.GET("/products", GetAll)

	//ruta para buscar por id
	router.GET("/products/:id", GetById)

	//corremos el servidor
	router.Run()
}