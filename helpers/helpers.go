package helpers

import (
	"errors"
	"finalProject/payment"
	"finalProject/product"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDatabase() (*gorm.DB, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-introduction?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, err := gorm.Open(mysql.Open(dbURL))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestIdValidity(id int64) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}

	return nil
}
func TestPriceValidity(price float32) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if price > 3.4e+38 {
		return errors.New("price must be less than infinity")
	}
	return nil
}

func PaymentToString(payment *payment.Payment, product *product.Product) string {
	return "Id: " + fmt.Sprintf("%v", payment.Id) + fmt.Sprintf("ProductId: %v", product.Id) + fmt.Sprintf("PricePaid: %v", payment.PricePaid) + fmt.Sprintf(" CreatedAt: %v", payment.CreatedAt) + fmt.Sprintf("UpdatedAt: %v", payment.UpdatedAt)
}
