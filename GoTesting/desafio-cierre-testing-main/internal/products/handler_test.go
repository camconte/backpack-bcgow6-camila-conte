package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
//estos son test de integracion
func createServer(mockRepository MockRepository) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)

	service := NewService(&mockRepository)
	handler := NewHandler(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("", handler.GetProducts)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	//el httptest.NewRecorder() guarda la response que obtiene el servidor
	return req, httptest.NewRecorder()
}

func TestGetProductsSuccess(t *testing.T) {
	//arrange
	database := []Product{
		{
			ID: "stub",
			SellerID: "1A",
			Description: "New Product",
			Price: 501.31,
		},
		{
			ID: "mock",
			SellerID: "5H0",
			Description: "limited edition product",
			Price: 300.1,
		},
		{
			ID: "dummy",
			SellerID: "1A",
			Description: "dummy double",
			Price: 100,
		},
	}

	mockRepository := MockRepository{
		Data: database,
	}

	//server creation
	r := createServer(mockRepository)

	//request creation
	request, responseRecorder := createRequestTest(http.MethodGet, "/products?seller_id=1A", "")

	//act
	r.ServeHTTP(responseRecorder, request)

	//assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, mockRepository.Data, responseRecorder.Body)
}

func TestGetProductsFail(t *testing.T) {
	//arrange
	database := []Product{
		{
			ID: "stub",
			SellerID: "1A",
			Description: "New Product",
			Price: 501.31,
		},
		{
			ID: "mock",
			SellerID: "5H0",
			Description: "limited edition product",
			Price: 300.1,
		},
		{
			ID: "dummy",
			SellerID: "1A",
			Description: "dummy double",
			Price: 100,
		},
	}

	mockRepository := MockRepository{
		Data: database,
	}

	//server creation
	r := createServer(mockRepository)

	//400 bad request creation - send no seller_id
	request400, responseRecorder400 := createRequestTest(http.MethodGet, "/products", "")

	//500 internal server error request creation - send a non-existent seller_id 
	request500, responseRecorder500 := createRequestTest(http.MethodGet, "/products?seller_id=aaa", "")

	//act
	//case 400 code
	r.ServeHTTP(responseRecorder400, request400)

	//case 500 code
	r.ServeHTTP(responseRecorder500, request500)

	//assert
	//bad request check
	assert.Equal(t, http.StatusBadRequest, responseRecorder400.Code)

	//internal server error check
	assert.Equal(t, http.StatusInternalServerError, responseRecorder500.Code)
}

