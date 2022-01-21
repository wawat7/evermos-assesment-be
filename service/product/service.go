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
	Update(product Product) Product
	GetPriceProduct(product Product) uint
	CreatePromotion(product Product, input CreateProductRequest) ProductPromotion
	UpdatePromotion(promotion ProductPromotion) ProductPromotion
	Delete(product Product)
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
		Stock:       int(input.Stock),
		TotalSold:   0,
		Rate:        0,
		Image:       input.Image,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	product = s.repository.Create(product)

	if input.Type != TYPE_NORMAL {
		_ = s.CreatePromotion(product, input)
	}
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

func (s *service) GetPriceProduct(product Product) uint {
	price := product.Price
	if product.Promotion.Id != 0 {
		price = product.Promotion.PriceAfterDiscount
	}
	return price
}

func (s *service) Update(product Product) Product {
	product = s.repository.Update(product)
	return product
}

func (s *service) CreatePromotion(product Product, input CreateProductRequest) ProductPromotion {
	promotion := ProductPromotion{
		ProductId:          product.Id,
		Type:               input.Type,
		IsActive:           true,
		Price:              product.Price,
		PriceAfterDiscount: input.PriceAfterDiscount,
		Stock:              int(input.StockPromotion),
		TotalSold:          0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	promotion = s.repository.CreatePromotion(promotion)
	return promotion
}

func (s *service) UpdatePromotion(promotion ProductPromotion) ProductPromotion {
	promotion = s.repository.UpdatePromotion(promotion)
	return promotion
}

func (s *service) Delete(product Product) {
	s.repository.Delete(product)

	return
}
