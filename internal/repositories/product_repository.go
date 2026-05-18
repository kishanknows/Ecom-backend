package repositories

import (
	"github.com/kishanknows/Ecom-backend/internal/database"
	"github.com/kishanknows/Ecom-backend/internal/models"
)

type ProductRepository struct {

}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetAll() ([]models.Product, error){
	var products []models.Product

	result := database.DB.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *ProductRepository) Create(product *models.Product) error {
	result := database.DB.Create(product)

	return result.Error
}