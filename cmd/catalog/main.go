package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/IgorShirokov/marketplace/internal/catalog/api"
	"github.com/IgorShirokov/marketplace/internal/catalog/api/handlers"
	"github.com/IgorShirokov/marketplace/internal/catalog/application/queries"
	"github.com/IgorShirokov/marketplace/internal/catalog/infrastructure/persistence"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	const (
		appPort = "9001"
		pgHost  = "localhost"
		pgPort  = "9101"
		pdDB    = "catalog-db-dev"
		pgUser  = "postgres"
		pgPass  = "12345678"
		pgSSL   = "disable"
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgHost, pgPort, pgUser, pgPass, pdDB, pgSSL,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	brandRepo := persistence.NewBrandRepository(db)
	listBrandsHandler := queries.NewBrandsHandler(brandRepo)
	brandsHandler := handlers.NewBrandsHandler(listBrandsHandler)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api.RegisterRoutes(r, brandsHandler)

	if err := r.Run(":" + appPort); err != nil {
		panic(err)
	}
}
