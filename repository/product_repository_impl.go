package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muhhylmi/store-api/helper"
	"github.com/muhhylmi/store-api/model/domain"
)

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product {
	SQL := "insert into Product(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, Product.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	Product.Id = int(id)

	return Product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product {
	SQL := "update Product set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, Product.Name, Product.Id)
	helper.PanicIfError(err)

	return Product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, Product domain.Product) {
	SQL := "delete from Product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, Product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, ProductId int) (domain.Product, error) {
	SQL := "select id, name from Product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, ProductId)
	helper.PanicIfError(err)
	defer rows.Close()

	Product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&Product.Id, &Product.Name)
		helper.PanicIfError(err)
		return Product, nil
	} else {
		return Product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "select id, name from Product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Product
	for rows.Next() {
		Product := domain.Product{}
		rows.Scan(&Product.Id, &Product.Name)
		categories = append(categories, Product)
	}

	return categories
}
