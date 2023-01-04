package payment

import "time"

type Payment struct {
	Id        int64     `json:"id"`
	ProductId int64     `json:"product_id"`
	PricePaid float32   `json:"price_paid"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
