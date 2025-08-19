package database

import (
	"context"
	"github.com/cynx-io/micro-name/internal/model/entity"
	"gorm.io/gorm"
)

type ExampleRepo struct {
	DB *gorm.DB
}

func NewExampleRepo(DB *gorm.DB) *ExampleRepo {
	return &ExampleRepo{
		DB: DB,
	}
}

func (r *ExampleRepo) GetExample(ctx context.Context, id int32) (*entity.TblExample, error) {
	return nil, nil
}
