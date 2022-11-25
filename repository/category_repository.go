package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/helmimuzkr/golang-restapi/model"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, category *model.Category) error
	GetAllCategory(ctx context.Context) ([]*model.Category, error)
	GetCategory(ctx context.Context, category *model.Category) (*model.Category, error)
}

type categoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepository {
	return &categoryRepository{DB: db}
}

// Create category
func (repository *categoryRepository) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("sql: start database transaction - %v", err)
	}

	query := "INSERT INTO categories(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("sql: insert category - %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("sql: commit transaction - %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("sql: getting last insert id - %v", err)
	}

	category.ID = int(id)

	return category, nil
}

// Update category
func (repository *categoryRepository) UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("sql: start database transaction - %v", err)
	}

	query := "UPDATE categories SET name = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, query, category.Name, category.ID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("sql: update category - %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("sql: commit transaction - %v", err)
	}

	return category, nil
}

// Delete category
func (repository *categoryRepository) DeleteCategory(ctx context.Context, category *model.Category) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return fmt.Errorf("sql: start database transaction - %v", err)
	}

	query := "DELETE FROM categories WHERE id = ?"
	_, err = tx.ExecContext(ctx, query, category.ID)
	if err != nil {
		return fmt.Errorf("sql: delete category - %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("sql: commit transaction - %v", err)
	}

	return nil
}

// Read Category By Id
func (repository *categoryRepository) GetCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	query := "SELECT * FROM categories WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, query, category.ID)
	if err != nil {
		return nil, fmt.Errorf("sql: select category by id - %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	err = rows.Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	return category, nil
}

// Read all category
func (repository *categoryRepository) GetAllCategory(ctx context.Context) ([]*model.Category, error) {
	query := "SELECT * FROM categories"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("sql: select category - %v", err)
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		category := new(model.Category)
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, fmt.Errorf("sql: no rows in result set")
		}

		categories = append(categories, category)
	}

	return categories, nil
}
