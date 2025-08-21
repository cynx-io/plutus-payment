package database

import (
	"context"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"gorm.io/gorm"
)

type TblCustomer struct {
	DB *gorm.DB
}

func NewCustomerRepo(DB *gorm.DB) *TblCustomer {
	return &TblCustomer{
		DB: DB,
	}
}

func (r *TblCustomer) GetCustomerByUserId(ctx context.Context, userId int32) (*entity.TblCustomer, error) {

	var existingCustomer entity.TblCustomer

	err := r.DB.WithContext(ctx).Where("user_id = ?", userId).First(&existingCustomer).Error
	if err != nil {
		// If there is another error, return it
		return nil, err
	}

	return &existingCustomer, nil
}

func (r *TblCustomer) CreateCustomer(ctx context.Context, customer entity.TblCustomer) (*entity.TblCustomer, error) {
	err := r.DB.WithContext(ctx).Create(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
