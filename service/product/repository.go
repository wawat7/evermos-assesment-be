package product

import (
	"evermos-assessment-be/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Create(product Product) Product
	FindAll() (products []Product)
	FindById(Id int) (product Product)
	Update(product Product) Product
	Delete(product Product)
	CreatePromotion(promotion ProductPromotion) ProductPromotion
	UpdatePromotion(promotion ProductPromotion) ProductPromotion
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) Create(product Product) Product {
	err := repo.db.Create(&product).Error
	helper.PanicIfError(err)

	return product
}

func (repo *repository) FindAll() (products []Product) {
	err := repo.db.Preload("Promotion").Find(&products).Error
	helper.PanicIfError(err)

	return
}

func (repo *repository) FindById(Id int) (product Product) {
	err := repo.db.Preload("Promotion", "product_promotions.is_active = true").Where("id = ?", Id).Find(&product).Error
	helper.PanicIfError(err)

	return
}

func (repo *repository) Update(product Product) Product {
	err := repo.db.Save(&product).Error
	helper.PanicIfError(err)

	return product
}

func (repo *repository) Delete(product Product) {
	err := repo.db.Delete(&product).Error
	helper.PanicIfError(err)

	return
}

func (repo *repository) CreatePromotion(promotion ProductPromotion) ProductPromotion {
	err := repo.db.Create(&promotion).Error
	helper.PanicIfError(err)

	return promotion
}

func (repo *repository) UpdatePromotion(promotion ProductPromotion) ProductPromotion {
	err := repo.db.Save(&promotion).Error
	helper.PanicIfError(err)

	return promotion
}
