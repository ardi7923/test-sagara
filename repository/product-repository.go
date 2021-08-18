package repository

import (
	"github.com/ardi7923/test-sagara/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	AllProduct() []entity.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) AllProduct() []entity.Product {
	var products []entity.Product
	db.connection.Find(&products)
	return products
}
