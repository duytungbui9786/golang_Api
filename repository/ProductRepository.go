package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type ProductRepo struct {
	Products map[int64]*model.Product
	autoID   int64
}

var Product ProductRepo

func init() {
	Product = ProductRepo{autoID: 0}
	Product.Products = make(map[int64]*model.Product)
	Product.InitData("sql:45312")
}

func (r *ProductRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ProductRepo) CreateNewProduct(Product *model.Product) int64 {
	nextID := r.getAutoID()
	Product.Id = nextID
	r.Products[nextID] = Product
	return nextID
}
func (r *ProductRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewProduct(&model.Product{
		CategoryId: 1,
		Name:       "ghế",
		Price:      50,
		Image:      "/uploads/images/item-02.jpg",
		IsSale:     true,
		Rating:     0,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000})

	r.CreateNewProduct(&model.Product{
		CategoryId: 1,
		Name:       "Denim jacket blue",
		Price:      92.5,
		Image:      "/uploads/images/item-03.jpg",
		IsSale:     false,
		Rating:     0,
		CreatedAt:  1610281342000,
		ModifiedAt: 1619283693000})
}

func (r *ProductRepo) GetAllProducts() map[int64]*model.Product {
	return r.Products
}

func (r *ProductRepo) FindProductById(Id int64) (*model.Product, error) {
	if Product, ok := r.Products[Id]; ok {
		return Product, nil
	} else {
		return nil, errors.New("Product not found")
	}
}
func (r *ProductRepo) GetPriceProductById(Id int64) float32 {
	if Product, ok := r.Products[Id]; ok {
		return Product.Price
	}
	return 0
}

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

func (r *ProductRepo) UpdateProduct(Product *model.Product) error {
	if _, ok := r.Products[Product.Id]; ok {
		r.Products[Product.Id] = Product
		return nil //tìm được
	} else {
		return errors.New("Product not found")
	}
}
