package repository

import (
	"errors"

	"github.com/TechMaster/golang/08Fiber/Basic/models"
)

type ProductRepo struct {
	Products map[int64]*models.Product
	autoID   int64
}

var Product ProductRepo

func (r *ProductRepo) GetAllProducts() map[int64]*models.Product {
	return r.Products
}

func (r *ProductRepo) FindProductById(Id int64) (*models.Product, error) {
	if Product, ok := r.Products[Id]; ok {
		return Product, nil
	} else {
		return nil, errors.New("Product not found")
	}
}
func (r *ProductRepo) GetPriceProductById(Id int64) float32 {
	if Product, ok := r.Products[Id]; ok {
		return Product.Newprice
	}
	return 0
}

// func (r *ProductRepo) CreateNewProduct(Product *models.Product) error {

// }

//del
func (r *ProductRepo) DeleteProductById(id int64) error {
	if _, ok := r.Products[id]; ok {
		delete(r.Products, id)
		return nil
	} else {
		return errors.New("Product not found")
	}
}

//create

func (r *ProductRepo) UpdateProduct(Product *models.Product) error {
	if _, ok := r.Products[Product.Id]; ok {
		r.Products[Product.Id] = Product
		return nil //tìm được
	} else {
		return errors.New("Product not found")
	}
}
