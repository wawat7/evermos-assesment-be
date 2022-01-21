package order

import (
	"evermos-assessment-be/helper"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	WithTrx(trxHandle *gorm.DB) *repository
	Create(order Order) Order
	FindAll() (orders []Order)
	FindById(Id int) (order Order)
	FindByProductId(productId int) (orders []Order)
	CreateOrderProduct(product OrderProduct) OrderProduct
	CreateOrderHistory(history OrderHistory) OrderHistory
	Delete(order Order)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) WithTrx(trxHandle *gorm.DB) *repository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return repo
	}
	repo.db = trxHandle
	return repo
}

func (repo *repository) Create(order Order) Order {
	err := repo.db.Create(&order).Error
	helper.PanicIfError(err)

	return order
}

func (repo *repository) FindAll() (orders []Order) {
	err := repo.db.Preload("OrderProduct").Preload("OrderHistory").Find(&orders).Error
	helper.PanicIfError(err)

	return
}

func (repo *repository) FindById(Id int) (order Order) {
	err := repo.db.Preload("OrderProduct").Preload("OrderHistory").Where("id = ?", Id).Find(&order).Error
	helper.PanicIfError(err)

	return
}

func (repo *repository) FindByProductId(productId int) (orders []Order) {
	err := repo.db.Preload("OrderProduct", "order_products.product_id = ?", productId).Find(&orders).Error
	helper.PanicIfError(err)

	return
}

func (repo *repository) CreateOrderProduct(product OrderProduct) OrderProduct {
	err := repo.db.Create(&product).Error
	helper.PanicIfError(err)

	return product
}

func (repo *repository) CreateOrderHistory(history OrderHistory) OrderHistory {
	err := repo.db.Create(&history).Error
	helper.PanicIfError(err)

	return history
}

func (repo *repository) Delete(order Order) {
	err := repo.db.Delete(&order).Error
	helper.PanicIfError(err)

	return
}
