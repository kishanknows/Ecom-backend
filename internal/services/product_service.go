package services

import (
	"github.com/kishanknows/Ecom-backend/internal/models"
	"github.com/kishanknows/Ecom-backend/internal/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		repository: repositories.NewProductRepository(),
	}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.repository.GetAll()
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repository.Create(product)
}