package database

import (
	"context"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"gorm.io/gorm"
)

type TblTokenInvoice struct {
	DB *gorm.DB
}

func NewTokenInvoiceRepo(DB *gorm.DB) *TblTokenInvoice {
	return &TblTokenInvoice{
		DB: DB,
	}
}

func (T *TblTokenInvoice) GetTokenInvoiceById(ctx context.Context, id string) (*entity.TblTokenInvoice, error) {
	var invoice entity.TblTokenInvoice
	err := T.DB.WithContext(ctx).Where("id = ?", id).First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (T *TblTokenInvoice) CreateTokenInvoice(ctx context.Context, invoice *entity.TblTokenInvoice) error {
	if err := T.DB.WithContext(ctx).Create(invoice).Error; err != nil {
		return err
	}
	return nil
}
