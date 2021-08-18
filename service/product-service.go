package service

import (
	"github.com/ardi7923/test-sagara/entity"
	"github.com/ardi7923/test-sagara/repository"
)

type ProductService interface {
	All() []entity.Product
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}
func (service *productService) All() []entity.Product {
	return service.productRepository.AllProduct()
}
