package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type ImageProductRepo struct {
	ImageProducts map[int64]*model.ImageProduct
	autoID        int64
}

var ImageProduct ImageProductRepo

func init() {
	ImageProduct = ImageProductRepo{autoID: 0}
	ImageProduct.ImageProducts = make(map[int64]*model.ImageProduct)
	ImageProduct.InitData("sql:45312")
}

func (r *ImageProductRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ImageProductRepo) CreateNewImageProduct(ImageProduct *model.ImageProduct) int64 {
	nextID := r.getAutoID()
	ImageProduct.Id = nextID
	r.ImageProducts[nextID] = ImageProduct
	return nextID
}
func (r *ImageProductRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewImageProduct(&model.ImageProduct{
		Id:        1,
		IdProduct: 1,
		Image:     "/uploads/images/item-04.jpg",
	})

	r.CreateNewImageProduct(&model.ImageProduct{
		Id:        2,
		IdProduct: 1,
		Image:     "/uploads/images/item-04.jpg",
	})
	r.CreateNewImageProduct(&model.ImageProduct{
		Id:        3,
		IdProduct: 1,
		Image:     "/uploads/images/item-04.jpg",
	})
}

func (r *ImageProductRepo) GetAllImageProducts() map[int64]*model.ImageProduct {
	return r.ImageProducts
}

//del
func (r *ImageProductRepo) DeleteImageProductById(id int64) error {
	if _, ok := r.ImageProducts[id]; ok {
		delete(r.ImageProducts, id)
		return nil
	} else {
		return errors.New("ImageProduct not found")
	}
}

//create
