package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type CategoryRepo struct {
	Categorys map[int64]*model.Category
	autoID    int64
}

var Category CategoryRepo

func init() {
	Category = CategoryRepo{autoID: 0}
	Category.Categorys = make(map[int64]*model.Category)
	Category.InitData("sql:45312")
}

func (r *CategoryRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *CategoryRepo) CreateNewCategory(Category *model.Category) int64 {
	nextID := r.getAutoID()
	Category.Id = nextID
	r.Categorys[nextID] = Category
	return nextID
}
func (r *CategoryRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewCategory(&model.Category{
		Id:   1,
		Name: "Women",
	})

	r.CreateNewCategory(&model.Category{
		Id:   2,
		Name: "Man",
	})
	r.CreateNewCategory(&model.Category{
		Id:   3,
		Name: "kid",
	})
}

func (r *CategoryRepo) GetAllCategorys() map[int64]*model.Category {
	return r.Categorys
}

func (r *CategoryRepo) FindCategoryById(Id int64) (*model.Category, error) {
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
	for _, v := range r.Products {
		if v.CategoryId == id {
			return errors.New("Product is found")
		}
	}
	return nil
}
func (r *CategoryRepo) UpdateCategory(Category *model.Category) error {
	if _, ok := r.Categorys[Category.Id]; ok {
		r.Categorys[Category.Id] = Category
		return nil //tìm được
	} else {
		return errors.New("Category not found")
	}
}
