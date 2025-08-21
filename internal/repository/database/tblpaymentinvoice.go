package database

import (
	"context"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"gorm.io/gorm"
)

type TblPaymentInvoice struct {
	DB *gorm.DB
}

func NewPaymentInvoiceRepo(DB *gorm.DB) *TblPaymentInvoice {
	return &TblPaymentInvoice{
		DB: DB,
	}
}

func (T *TblPaymentInvoice) GetPaymentInvoiceById(ctx context.Context, id string) (*entity.TblPaymentInvoice, error) {
	var invoice entity.TblPaymentInvoice
	err := T.DB.WithContext(ctx).Where("id = ?", id).First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (T *TblPaymentInvoice) CreatePaymentInvoice(ctx context.Context, invoice *entity.TblPaymentInvoice) error {
	if err := T.DB.WithContext(ctx).Create(invoice).Error; err != nil {
		return err
	}
	return nil
}
