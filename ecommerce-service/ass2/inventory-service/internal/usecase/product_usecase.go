package usecase

import (
	"fmt"
	"inventory-service/internal/entity"
	"inventory-service/internal/repository"
)

type ProductUsecase struct {
	repo *repository.ProductRepository
}

func NewProductUsecase(repo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) GetProduct(id string) (*entity.Product, error) {
	return uc.repo.GetProductByID(id)
}

func (uc *ProductUsecase) CreateProduct(product entity.Product) (entity.Product, error) {
	if product.Price < 0 {
		return entity.Product{}, fmt.Errorf("price cannot be negative")
	}

	err := uc.repo.CreateProduct(&product)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (uc *ProductUsecase) UpdateProduct(product *entity.Product) error {
	return uc.repo.UpdateProduct(product)
}

func (uc *ProductUsecase) DeleteProduct(id string) error {
	return uc.repo.DeleteProduct(id)
}

func (uc *ProductUsecase) GetAllProducts(name string, category string, limit, offset int) ([]entity.Product, error) {
	return uc.repo.GetAllProducts(name, category, limit, offset)
}
