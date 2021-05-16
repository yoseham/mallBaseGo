package services

import (
	"app/datamodels"
	"app/repositories"
)

type IproductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) (bool, error)
	InsertProduct(*datamodels.Product) (int64, error)
	UpdateProduct(*datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

func NewProductService(product repositories.IProduct) IproductService {
	return &ProductService{product}
}

func (p *ProductService) GetProductByID(productID int64) (*datamodels.Product, error) {
	return p.productRepository.SelectByKey(productID)
}

func (p *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}

func (p *ProductService) DeleteProductByID(productID int64) (bool, error) {
	return p.productRepository.Delete(productID)
}

func (p *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}
