package repository

import (
	"errors"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
}
func (r *ReviewRepo) GetAllReview() map[int64]*model.Review {
	return r.reviews
}
func (r *ReviewRepo) getAutoIDRe() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoIDRe()
	review.Id = nextID
	r.reviews[nextID] = review
	return nextID
}

func (r *ProductRepo) CheckProduct(review *model.Review) error {
	if _, ok := r.Products[review.ProductId]; ok {
		return nil //tìm được
	} else {
		return errors.New("Product not found")
	}
}

func (r *ReviewRepo) AverageRating(id int64) (result map[int64]float32) {
	sum := make(map[int64]int)
	reviewProduct := make(map[int64]*model.Review)
	number := make(map[int64]int)
	result = make(map[int64]float32)

	for _, value := range r.reviews {
		if value.ProductId == int64(id) {
			reviewProduct[value.Id] = value
		}
	}
	// database.DB.Where("BookId", review.BookId)
	// arr := reviewBook[1]
	for _, value := range reviewProduct {
		number[value.ProductId]++
		sum[value.ProductId] += value.Rating
	}
	for key := range number {
		result[key] = float32(sum[key]) / float32(number[key])
	}
	return result
}
func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}
