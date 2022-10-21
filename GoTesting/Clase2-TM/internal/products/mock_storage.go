package products

type MockStorage struct{
	dataMock []Product
	ReadMethodWasCall bool
	errOnRead error
	errOnWrite error
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errOnWrite != nil{
		return m.errOnWrite
	}

	dataProducts := data.([]Product)
	m.dataMock = dataProducts
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errOnRead != nil {
		return m.errOnRead
	}
	
	dataProducts := data.(*[]Product)
	*dataProducts = m.dataMock
	m.ReadMethodWasCall = true
	return nil
}

