package service

import (
	"github.com/comerc/go-clean-architecture/internal/model"
)

type IProductRepository interface {
	Insert(*model.ProductModel) (*model.ProductModel, error)
	Update(*model.ProductModel) error
	Delete(id string) error
	FindById(id string) (*model.ProductModel, error)
	Find(limit, offset int, where any) ([]model.ProductModel, error)
}

type ProductService struct {
	repository IProductRepository
}

func New(r IProductRepository) *ProductService {
	return &ProductService{
		repository: r,
	}
}

func (s *ProductService) Create(model *model.ProductModel) (*model.ProductModel, error) {
	return s.repository.Insert(model)
}

func (s *ProductService) Update(model *model.ProductModel) error {
	return s.repository.Update(model)
}

func (s *ProductService) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *ProductService) ReadById(id string) (*model.ProductModel, error) {
	return s.repository.FindById(id)
}

func (s *ProductService) Read(limit, offset int, where any) ([]model.ProductModel, error) {
	return s.repository.Find(limit, offset, where)
}
