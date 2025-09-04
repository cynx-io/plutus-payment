package database

import (
	"context"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"gorm.io/gorm"
)

type TblTokenPriceList struct {
	DB *gorm.DB
}

func NewTokenPriceListRepo(DB *gorm.DB) *TblTokenPriceList {
	return &TblTokenPriceList{
		DB: DB,
	}
}

func (r *TblTokenPriceList) GetTokenPriceListById(ctx context.Context, id int32) (*entity.TblTokenPriceList, error) {

	var existingTokenPriceList entity.TblTokenPriceList

	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existingTokenPriceList).Error
	if err != nil {
		return nil, err
	}

	return &existingTokenPriceList, nil
}

func (r *TblTokenPriceList) ListTokenPriceList(ctx context.Context) ([]*entity.TblTokenPriceList, error) {
	var tokenPriceLists []*entity.TblTokenPriceList

	err := r.DB.WithContext(ctx).Find(&tokenPriceLists).Error
	if err != nil {
		return nil, err
	}

	return tokenPriceLists, nil
}
