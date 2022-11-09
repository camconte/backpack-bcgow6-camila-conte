package domain

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ProductType string `json:"productType"`
	Count int `json:"count"`
	Price float64 `json:"price"`
	WarehouseId int `json:"warehouse_id"`
}