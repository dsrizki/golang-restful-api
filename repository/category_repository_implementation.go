package repository

import (
	"context"
	"database/sql"
	"dsrizki/golang-restful-api/helper"
	"dsrizki/golang-restful-api/model/domain"
	"errors"
)

type CategoryRepositoryImplementation struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImplementation{}
}

func (repository *CategoryRepositoryImplementation) Save(context context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values (?)"

	result, err := tx.ExecContext(context, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImplementation) Update(context context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(context, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImplementation) Delete(context context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(context, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImplementation) FindById(context context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"

	rows, err := tx.QueryContext(context, SQL, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("Category is not found")
	}
}

func (repository *CategoryRepositoryImplementation) FindAll(context context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"

	rows, err := tx.QueryContext(context, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}

		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}
	return categories
}
