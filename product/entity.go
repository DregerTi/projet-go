package product

import "time"

type Product struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
