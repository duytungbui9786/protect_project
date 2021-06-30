package repository

import (
	"errors"

	"github.com/TechMaster/golang/08Fiber/Basic/models"
)

type CategoryRepo struct {
	Categorys map[int64]*models.Category
	autoID    int64
}

var Category CategoryRepo

func (r *CategoryRepo) GetAllCategorys() map[int64]*models.Category {
	return r.Categorys
}

func (r *CategoryRepo) FindCategoryById(Id int64) (*models.Category, error) {
	if Category, ok := r.Categorys[Id]; ok {
		return Category, nil
	} else {
		return nil, errors.New("Category not found")
	}
}

//del
func (r *CategoryRepo) DeleteCategoryById(id int64) error {
	if _, ok := r.Categorys[id]; ok {
		delete(r.Categorys, id)
		return nil
	} else {
		return errors.New("Category not found")
	}
}

//create
func (r *ProductRepo) CheckProductUseCategory(id int64) error {
	// for _, v := range r.Products {
	// 	if v.CategoryId == id {
	// 		return errors.New("Product is found")
	// 	}
	// }
	return nil
}
func (r *CategoryRepo) UpdateCategory(Category *models.Category) error {
	if _, ok := r.Categorys[Category.Id]; ok {
		r.Categorys[Category.Id] = Category
		return nil //tìm được
	} else {
		return errors.New("Category not found")
	}
}
