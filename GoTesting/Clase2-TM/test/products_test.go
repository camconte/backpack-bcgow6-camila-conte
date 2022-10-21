package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/cmd/server/handler"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/internal/products"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore mocks.MockStorage) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456clc")

	gin.SetMode(gin.ReleaseMode)

	repo := products.NewRepository(&mockStore)
	service := products.NewService(repo)
	pHandler := handler.NewProductHandler(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.PUT("/:id", pHandler.Update())
	pr.DELETE("/:id", pHandler.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456clc")

	return req, httptest.NewRecorder()
}

func TestUpdateOk(t *testing.T) {
	//arrange
	dataMockStorage := []products.Product{
		{
			Id:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Colour: "Blue",
			Stock: 2000,
			Price: 300,
			Code: "1234",
			Published: true,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
		{
			Id:		2,
			Name:  "Rexona",
			Colour: "Pink",
			Stock: 34,
			Price: 100,
			Code: "456",
			Published: false,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
	}
	
	mockStorage := mocks.MockStorage{
		DataMock: dataMockStorage,
	}

	//creamos el server y definimos las rutas
	r := createServer(mockStorage)

	//creamos el request de tipo PUT y le pasamos la ruta y el json con la data
	request, responseRecorder := createRequestTest(http.MethodPut, "/products/1", `{
        "name": "Tester",
		"colour": "Blue",
		"stock": 10,
		"price": 90,
		"code": "12345A",
		"published": false
    }`)
	
	//act
	//indicamos al servidor que atienda la solicitud
	r.ServeHTTP(responseRecorder, request)

	//assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestDeleteOk(t *testing.T) {
	//arrange
	dataMockStorage := []products.Product{
		{
			Id:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Colour: "Blue",
			Stock: 2000,
			Price: 300,
			Code: "1234",
			Published: true,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
		{
			Id:		2,
			Name:  "Rexona",
			Colour: "Pink",
			Stock: 34,
			Price: 100,
			Code: "456",
			Published: false,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
	}
	
	mockStorage := mocks.MockStorage{
		DataMock: dataMockStorage,
	}

	//creamos el server y definimos las rutas
	r := createServer(mockStorage)

	request, responseRecorder := createRequestTest(http.MethodDelete, "/products/1", "")
	
	//act

	r.ServeHTTP(responseRecorder, request)

	//assert
	
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}