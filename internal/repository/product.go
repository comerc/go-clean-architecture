package repository

import (
	"github.com/comerc/go-clean-architecture/internal/model"
	"github.com/comerc/go-clean-architecture/internal/service"
)

type ProductRepository struct {
	service.IProductRepository
}

func (r *ProductRepository) Insert(*model.ProductModel) (*model.ProductModel, error) {
	return nil, nil
}

func (r *ProductRepository) Update(*model.ProductModel) error {
	return nil
}

func (r *ProductRepository) Delete(id string) error {
	return nil
}

func (r *ProductRepository) FindById(id string) (*model.ProductModel, error) {
	return nil, nil
}

func (r *ProductRepository) Find(limit, offset int, where any) ([]model.ProductModel, error) {
	return nil, nil
}
