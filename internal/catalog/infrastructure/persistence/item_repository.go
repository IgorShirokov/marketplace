package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/IgorShirokov/marketplace/internal/catalog/domain/entities"
	"github.com/google/uuid"
)

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *itemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Items(ctx context.Context) ([]entities.CatalogItem, error) {
	query := `
		SELECT 
		ci.id, ci.title, ci.short_description, ci.full_description, ci.image_url, ci.price,
		b.id, b.title, c.id, c.title
		FROM catalog_items ci
		LEFT JOIN brands b ON ci.brand_id = b.id
		LEFT JOIN categories c ON ci.category_id = c.id 
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Items query %w", err)
	}
	defer rows.Close()

	var items []entities.CatalogItem
	for rows.Next() {
		item, err := scanCatalogItem(rows)
		if err != nil {
			return nil, fmt.Errorf("scatn item %w", err)
		}
		items = append(items, item)
	}

	return items, rows.Err()

}

func scanCatalogItem(rows *sql.Rows) (entities.CatalogItem, error) {
	var item entities.CatalogItem
	var brandID *uuid.UUID
	var brandTitle *string
	var categoryID *uuid.UUID
	var categoryTitle *string

	err := rows.Scan(
		&item.ID,
		&item.Title,
		&item.ShortDescription,
		&item.FullDescription,
		&item.ImageURL,
		&item.Price,
		&brandID,
		&brandTitle,
		&categoryID,
		&categoryTitle,
	)

	if err != nil {
		return item, err
	}

	if brandID != nil {
		item.Brand = &entities.Brand{
			BaseEntity: entities.BaseEntity{
				ID:    brandID.String(),
				Title: *brandTitle,
			},
		}
	}

	if categoryID != nil {
		item.Category = &entities.Category{
			BaseEntity: entities.BaseEntity{
				ID:    categoryID.String(),
				Title: *categoryTitle,
			},
		}
	}

	return item, nil
}
