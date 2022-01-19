package user

import (
	"evermos-assessment-be/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) User
	FindAll() (users []User)
	FindById(Id int) (user User)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(user User) User {
	err := r.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *repository) FindAll() (users []User) {
	err := r.db.Find(&users).Error
	helper.PanicIfError(err)

	return
}

func (r *repository) FindById(Id int) (user User) {
	err := r.db.Where("id = ?", Id).Find(&user).Error
	helper.PanicIfError(err)

	return
}
