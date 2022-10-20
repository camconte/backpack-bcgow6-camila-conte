package products

type MockStorage struct{
	dataMock []Product
	ReadMethodWasCall bool
}

func (m *MockStorage) Write(data interface{}) error {
	dataProducts := data.([]Product)
	m.dataMock = dataProducts
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	dataProducts := data.(*[]Product)
	*dataProducts = m.dataMock
	m.ReadMethodWasCall = true
	return nil
}

