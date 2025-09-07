package database

import (
	"context"
	"errors"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TblBalance struct {
	DB *gorm.DB
}

func NewBalanceRepo(DB *gorm.DB) *TblBalance {
	return &TblBalance{
		DB: DB,
	}
}

func (T *TblBalance) GetBalanceByUserId(ctx context.Context, userId int32) (*entity.TblBalance, error) {
	var balance entity.TblBalance
	err := T.DB.WithContext(ctx).Where("user_id = ?", userId).First(&balance).Error
	if err == nil {
		return &balance, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = T.CreateBalance(ctx, &entity.TblBalance{
		UserId:       userId,
		TokenBalance: 0,
	})
	if err != nil {
		return nil, err
	}

	return T.GetBalanceByUserId(ctx, userId)
}

func (T *TblBalance) CreateBalance(ctx context.Context, balance *entity.TblBalance) error {
	if err := T.DB.WithContext(ctx).Create(balance).Error; err != nil {
		return err
	}
	return nil
}

func (T *TblBalance) IncrementBalance(ctx context.Context, userId int32, delta float32) error {
	return T.DB.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"token_balance": gorm.Expr("token_balance + ?", delta),
			}),
		}).
		Create(&entity.TblBalance{
			UserId:       userId,
			TokenBalance: delta,
		}).Error
}

func (T *TblBalance) DecrementBalance(ctx context.Context, userId int32, delta float32) error {
	result := T.DB.WithContext(ctx).
		Model(&entity.TblBalance{}).
		Where("user_id = ? AND token_balance >= ?", userId, delta).
		Update("token_balance", gorm.Expr("token_balance - ?", delta))

	if result.RowsAffected == 0 {
		return errors.New("insufficient balance or user not found")
	}

	return result.Error
}
