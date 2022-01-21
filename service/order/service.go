package order

import (
	"errors"
	"evermos-assessment-be/helper"
	"evermos-assessment-be/service/product"
	"evermos-assessment-be/service/user"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

type Service interface {
	WithTrx(trxHandle *gorm.DB) *service
	Save(input CreateOrderRequest) (Order, error)
	FindAll() (orders []Order)
	FindById(Id int) (order Order, err error)
	CalculateTotal(product product.Product, quantity int) uint
	CreateOrderProduct(order Order, productData product.Product, quantity int) OrderProduct
	CreateOrderHistory(order Order, userData user.User) OrderHistory
	UpdateStockAndTotalSoldProduct(product product.Product, quantity int)
	ValidationStockProduct(productData product.Product, quantity int) bool
}

type service struct {
	repository     Repository
	serviceUser    user.Service
	serviceProduct product.Service
}

func NewService(repository Repository, serviceUser user.Service, serviceProduct product.Service) *service {
	return &service{repository: repository, serviceUser: serviceUser, serviceProduct: serviceProduct}
}

func (s *service) WithTrx(trxHandle *gorm.DB) *service {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s *service) Save(input CreateOrderRequest) (Order, error) {
	var order Order

	userInfo, err := s.serviceUser.FindById(input.UserId)
	if err != nil {
		return order, err
	}
	userInfoString := helper.ConvertDataToJsonString(userInfo)

	productData, err := s.serviceProduct.FindById(input.ProductId)
	if err != nil {
		return order, err
	}
	total := s.CalculateTotal(productData, input.Quantity)

	order = Order{
		UserInfo:    userInfoString,
		UserId:      userInfo.Id,
		Code:        "TRX-" + strconv.Itoa(rand.Int()),
		Total:       int(total),
		Status:      STATUS_PENDING,
		Description: "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	order = s.repository.Create(order)

	_ = s.CreateOrderProduct(order, productData, input.Quantity)
	_ = s.CreateOrderHistory(order, userInfo)

	if s.ValidationStockProduct(productData, input.Quantity) == false {
		return order, errors.New("out of stock")
	}

	s.UpdateStockAndTotalSoldProduct(productData, input.Quantity)

	return order, nil
}

func (s *service) FindAll() (orders []Order) {
	orders = s.repository.FindAll()
	return
}

func (s *service) FindById(Id int) (order Order, err error) {
	order = s.repository.FindById(Id)
	if order.Id == 0 {
		return order, errors.New("order not found")
	}
	return order, nil
}

func (s *service) CreateOrderProduct(order Order, productData product.Product, quantity int) OrderProduct {
	productJson := helper.ConvertDataToJsonString(productData)
	subTotal := int(s.serviceProduct.GetPriceProduct(productData)) * quantity

	orderProduct := OrderProduct{
		OrderId:     order.Id,
		ProductInfo: productJson,
		ProductId:   productData.Id,
		Quantity:    quantity,
		SubTotal:    subTotal,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	orderProduct = s.repository.CreateOrderProduct(orderProduct)
	return orderProduct
}

func (s *service) CreateOrderHistory(order Order, userData user.User) OrderHistory {
	history := OrderHistory{
		OrderId:   order.Id,
		Status:    order.Status,
		CreatedBy: userData.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	history = s.repository.CreateOrderHistory(history)
	return history
}

func (s *service) CalculateTotal(product product.Product, quantity int) uint {
	var total uint

	total = s.serviceProduct.GetPriceProduct(product) * uint(quantity)
	return total
}

func (s *service) UpdateStockAndTotalSoldProduct(product product.Product, quantity int) {
	product = DecreaseStockProduct(product, quantity)
	product = IncreaseTotalSoldProduct(product, quantity)

	product = s.serviceProduct.Update(product)
	if product.Promotion.Id != 0 {
		_ = s.serviceProduct.UpdatePromotion(product.Promotion)
	}
}

func (s *service) ValidationStockProduct(productData product.Product, quantity int) bool {

	productStock := productData.Stock
	if productData.Promotion.Id != 0 {
		productStock = productData.Promotion.Stock
	}

	stock := int(productStock) - quantity
	if stock < 0 {
		return false
	}

	return true

}

func DecreaseStockProduct(product product.Product, quantity int) product.Product {

	if product.Promotion.Id != 0 {
		stockPromo := int(product.Promotion.Stock)
		stockPromo -= quantity
		if stockPromo < 0 {
			product.Promotion.Stock -= uint(quantity)
		} else {
			helper.PanicIfError(errors.New("out of stock"))
		}
	} else {
		stock := int(product.Stock)
		stock -= quantity
		if stock < 0 {
			product.Stock -= uint(quantity)
		} else {
			helper.PanicIfError(errors.New("out of stock"))
		}
	}

	return product
}

func IncreaseTotalSoldProduct(product product.Product, quantity int) product.Product {

	if product.Promotion.Id != 0 {
		product.Promotion.TotalSold += uint(quantity)
	}

	product.TotalSold += uint(quantity)
	return product
}
