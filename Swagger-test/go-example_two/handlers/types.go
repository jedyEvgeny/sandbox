package handlers

type Item struct {
	Product string  `json:"product" example:"Ананас"`
	Price   float64 `json:"price" example:"264.50"`
}

// Resourse для примера возвращаемых значений в OpenAPI, сгенерированной Swaggo
type Resourse struct {
	ID string `json:"id" example:"Ананас.txt"`
}

// ItemPrice для примера возвращаемых значений в OpenAPI, сгенерированной Swaggo
type ItemPrice struct {
	Price string `json:"price" example:"264.50"`
}

// Resourse для примера возвращаемых значений в OpenAPI, сгенерированной Swaggo
type ItemProduct struct {
	ID string `json:"id" example:"Ananas"`
}
