package repositories

import (
	"context"

	"github.com/IgorShirokov/marketplace/internal/catalog/domain/entities"
)

type BrandRepository interface {
	Brands(ctx context.Context) ([]entities.Brand, error)
}

type CategoryRepository interface {
	Categories(ctx context.Context) ([]entities.Category, error)
}

type CatalogItemRepository interface {
	Items(ctx context.Context) ([]entities.CatalogItem, error)
}
