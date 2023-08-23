package controller

import (
	"github.com/comerc/go-clean-architecture/internal/service"
)

type ProductController struct {
	service.ProductService // !!
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}
