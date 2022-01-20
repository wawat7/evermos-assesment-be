package product

import (
	"errors"
	"evermos-assessment-be/helper"
	"time"
)

type Service interface {
	Save(input CreateProductRequest) Product
	FindAll() (products []Product)
	FindById(Id int) (product Product, err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Save(input CreateProductRequest) Product {

	ingredients := helper.ConvertDataToJsonString(input.Ingredient)
	product := Product{
		Name:        input.Name,
		Description: input.Description,
		Ingredient:  ingredients,
		Price:       input.Price,
		Stock:       input.Stock,
		Rate:        0,
		Image:       input.Image,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	product = s.repository.Create(product)

	if input.Type != TYPE_NORMAL {
		promotion := ProductPromotion{
			ProductId:          product.Id,
			Type:               input.Type,
			IsActive:           true,
			Price:              product.Price,
			PriceAfterDiscount: input.PriceAfterDiscount,
			Stock:              input.StockPromotion,
			TotalSold:          0,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		promotion = s.repository.CreatePromotion(promotion)

	}

	product = s.repository.FindById(product.Id)
	return product
}

func (s *service) FindAll() (products []Product) {
	products = s.repository.FindAll()
	return
}

func (s *service) FindById(Id int) (product Product, err error) {
	product = s.repository.FindById(Id)
	if product.Id == 0 {
		return product, errors.New("product not found")
	}
	return
}
