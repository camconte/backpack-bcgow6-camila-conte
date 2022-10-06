package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/internal/products"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/pkg/web"
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

//ListProducts godoc
//@Summary List products
//@Tags Products
//@Description Get products from a file
//@Produce json
//@Param token header string true "Token"
//@Success 200 {object} web.Response
//@Failure 401 {object} web.Response
//@Failure 500 {object} web.Response
//@Router /products [GET]
func (h *ProductHandler) GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "you don't have permissions to make that request"))
			return
		}

		products, err := h.service.GetAll()
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
	}
}

//StoreProducts godoc
//@Summary Store products
//@Tags Products
//@Description Store a new product in a file
//@Accept json
//@Produce json
//@Param token header string true "Token"
//@Param product body request true "Product to store"
//@Success 200 {object} web.Response
//@Failure 400 {object} web.Response
//@Failure 401 {object} web.Response
//@Failure 500 {object} web.Response
//@Router /products [POST]
func (h *ProductHandler) Store() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "you don't have permissions to make that request"))
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil{
			
			if req.Name == "" {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the name is required"))
				return			}
			if req.Colour == "" {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the colour is required"))
				return			}
			if req.Price == 0 {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the price is required"))
				return			}
			if req.Stock == 0 {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the stock is required"))
				return			}
			if req.Code == "" {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the code is required"))
				return
			}

		}

		newProduct, err := h.service.Store(req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published)
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, newProduct, ""))

	}
}

//UpdateProduct godoc
//@Summary Update products
//@Tags Products
//@Description Update an entire product of the file
//@Accept json
//@Produce json
//@Param token header string true "Token"
//@Param id path int true "Product ID"
//@Param product body request true "Product to update"
//@Success 200 {object} web.Response
//@Failure 400 {object} web.Response
//@Failure 401 {object} web.Response
//@Failure 404 {object} web.Response
//@Router /products/{id} [PUT]
func (h *ProductHandler) Update() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "you don't have permissions to make that request"))
			return
		}

		//recolectamos el id que viene por path param
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil{
			
			if req.Name == "" {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the name is required"))
				return			}
			if req.Colour == "" {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the colour is required"))
				return			}
			if req.Price == 0 {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the price is required"))
				return			}
			if req.Stock == 0 {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the stock is required"))
				return			}
			if req.Code == "" {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "the code is required"))
				return
			}
		}

		updatedProduct, err := h.service.Update(int(id), req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedProduct, ""))


	}
}

//UpdateNameAndPrice godoc
//@Summary Update product's name and product's price
//@Tags Products
//@Description Update product's name, product's price or both
//@Accept json
//@Produce json
//@Param token header string true "Token"
//@Param id path int true "Product ID"
//@Param fieldsToUpdate body request true "Name, price or both to update"
//@Success 200 {object} web.Response
//@Failure 400 {object} web.Response
//@Failure 401 {object} web.Response
//@Failure 404 {object} web.Response
//@Router /products/{id} [PATCH]
func (h *ProductHandler) UpdateNameAndPrice() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "you don't have permissions to make that request"))
			return
		}

		//recolectamos el id que viene por path param
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil{
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if req.Name == "" && req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "name or price must be send to update the product"))
			return
		}

		updatedProduct, err := h.service.UpdateNameAndPrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedProduct, ""))
		
	}
}

//DeleteProduct godoc
//@Summary Delete product
//@Tags Products
//@Param token header string true "Token"
//@Param id path int true "Product ID"
//@Success 200 {object} web.Response
//@Failure 400 {object} web.Response
//@Failure 401 {object} web.Response
//@Failure 404 {object} web.Response
//@Router /products/{id} [DELETE]
func (h *ProductHandler) Delete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "you don't have permissions to make that request"))
			return
		}

		//recolectamos el id que viene por path param
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		
		err = h.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("The product with id %d has been removed", id), ""))

	}
}