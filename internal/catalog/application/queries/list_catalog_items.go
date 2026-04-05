package queries

import (
	"context"

	"github.com/IgorShirokov/marketplace/internal/catalog/domain/entities"
	"github.com/IgorShirokov/marketplace/internal/catalog/domain/repositories"
)

type CatalogItemsHandler struct {
	repo repositories.CatalogItemRepository
}

func NewCatalogItemsHandler(repo repositories.CatalogItemRepository) *CatalogItemsHandler {
	return &CatalogItemsHandler{repo: repo}
}

func (h *CatalogItemsHandler) Handle(ctx context.Context) ([]entities.CatalogItem, error) {
	return h.repo.Items(ctx)
}
