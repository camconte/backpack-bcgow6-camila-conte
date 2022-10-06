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

func (h *ProductHandler) GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "you don't have permissions to make that request"))
			return
		}

		products, err := h.service.GetAll()
		if err != nil{
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
	}
}

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
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, newProduct, ""))

	}
}

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