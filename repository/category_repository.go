package repository

import (
	"context"
	"database/sql"
	"dsrizki/golang-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(context context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(context context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(context context.Context, tx *sql.Tx, category domain.Category)
	FindById(context context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(context context.Context, tx *sql.Tx) []domain.Category
}
