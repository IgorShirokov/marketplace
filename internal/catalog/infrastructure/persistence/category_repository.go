package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/IgorShirokov/marketplace/internal/catalog/domain/entities"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Categories(ctx context.Context) ([]entities.Category, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, title FROM categories ORDER BY title`)
	if err != nil {
		return nil, fmt.Errorf("categories query: %w", err)
	}
	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var c entities.Category
		if err := rows.Scan(&c.ID, &c.Title); err != nil {
			return nil, fmt.Errorf("scan categories: %w", err)
		}
		categories = append(categories, c)
	}

	return categories, rows.Err()
}
