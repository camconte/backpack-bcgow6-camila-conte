package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct{
	Name string `json:"name"`
	Colour string `json:"colour"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Code string `json:"code"`
	Published bool `json:"published"`
}

//estructura del controlador
type ProductHandler struct{
	service products.Service
}

//retorna el controlador
func NewProductHandler(ps products.Service) *ProductHandler{
	return &ProductHandler{
		service: ps,
	}
}

func (h *ProductHandler) GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you don't have permissions to make that request",
			})
			return
		}

		products, err := h.service.GetAll()
		if err != nil{
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, products)
	}
}

func (h *ProductHandler) Store() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you don't have permissions to make that request",
			})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil{
			
			//recorremos los errores detallados y los almacenamos para luego mostrarlos todos juntos
			errorMessages := []string{}

			if req.Name == "" {
				errorMessages = append(errorMessages, "the name is required")
			}
			if req.Colour == "" {
				errorMessages = append(errorMessages, "the colour is required")
			}
			if req.Price == 0 {
				errorMessages = append(errorMessages, "the price is required")
			}
			if req.Stock == 0 {
				errorMessages = append(errorMessages, "the stock is required")
			}
			if req.Code == "" {
				errorMessages = append(errorMessages, "the code is required")
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})

			return
		}

		newProduct, err := h.service.Store(req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, newProduct)

	}
}

func (h *ProductHandler) Update() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you don't have permissions to make that request",
			})
			return
		}

		//recolectamos el id que viene por path param
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil{
			
			//recorremos los errores detallados y los almacenamos para luego mostrarlos todos juntos
			errorMessages := []string{}

			if req.Name == "" {
				errorMessages = append(errorMessages, "the name is required")
			}
			if req.Colour == "" {
				errorMessages = append(errorMessages, "the colour is required")
			}
			if req.Price == 0 {
				errorMessages = append(errorMessages, "the price is required")
			}
			if req.Stock == 0 {
				errorMessages = append(errorMessages, "the stock is required")
			}
			if req.Code == "" {
				errorMessages = append(errorMessages, "the code is required")
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}

		updatedProduct, err := h.service.Update(int(id), req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, updatedProduct)


	}
}

func (h *ProductHandler) UpdateNameAndPrice() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you don't have permissions to make that request",
			})
			return
		}

		//recolectamos el id que viene por path param
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error"	: err.Error(),
			})
			return
		}

		if req.Name == "" && req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "name or price must be send to update the product",
			})
			return
		}

		updatedProduct, err := h.service.UpdateNameAndPrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, updatedProduct)
		
	}
}

func (h *ProductHandler) Delete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you don't have permissions to make that request",
			})
			return
		}

		//recolectamos el id que viene por path param
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		
		err0 := h.service.Delete(int(id))
		if err0 != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err0.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"info": fmt.Sprintf("The product with id %d has been removed", id),
		})

	}
}