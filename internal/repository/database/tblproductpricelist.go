package database

import (
	"context"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"gorm.io/gorm"
)

type TblProductPriceList struct {
	DB *gorm.DB
}

func NewProductPriceListRepo(DB *gorm.DB) *TblProductPriceList {
	return &TblProductPriceList{
		DB: DB,
	}
}

func (r *TblProductPriceList) GetProductPriceListById(ctx context.Context, id string) (*entity.TblProductPriceList, error) {

	var existingProductPriceList entity.TblProductPriceList

	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existingProductPriceList).Error
	if err != nil {
		return nil, err
	}

	return &existingProductPriceList, nil
}
