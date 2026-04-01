package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/IgorShirokov/marketplace/internal/catalog/domain/entities"
)

type brandRepository struct {
	db *sql.DB
}

func NewBrandRepository(db *sql.DB) *brandRepository {
	return &brandRepository{db: db}
}

func (r *brandRepository) Brands(ctx context.Context) ([]entities.Brand, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, title FROM brands ORDER BY title`)
	if err != nil {
		return nil, fmt.Errorf("brands query: %w", err)
	}
	defer rows.Close()

	var brands []entities.Brand

	for rows.Next() {
		var b entities.Brand
		if err := rows.Scan(&b.ID, &b.Title); err != nil {
			return nil, fmt.Errorf("scan brands: %w", err)
		}
		brands = append(brands, b)
	}

	return brands, rows.Err()
}
