package mocks

import "github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/internal/products"

type MockStorage struct{
	DataMock []products.Product
	ReadMethodWasCall bool
	errOnRead error
	errOnWrite error
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errOnWrite != nil{
		return m.errOnWrite
	}

	dataProducts := data.([]products.Product)
	m.DataMock = dataProducts
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errOnRead != nil {
		return m.errOnRead
	}
	
	dataProducts := data.(*[]products.Product)
	*dataProducts = m.DataMock
	m.ReadMethodWasCall = true
	return nil
}
