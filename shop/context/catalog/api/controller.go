package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/app"
	"github.com/jmdavril/template/shop/context/catalog/domain"
	"github.com/lib/pq"
	"net/http"
)

var logger = app.LoggerWith(app.Catalog, app.Api)

type CatalogController struct {
	router         *gin.Engine
	productService *domain.ProductService
}

func NewCatalogController(r *gin.Engine, s *domain.ProductService) *CatalogController {
	return &CatalogController{
		router:         r,
		productService: s,
	}
}

func (c *CatalogController) handleProductCreate() gin.HandlerFunc {
	type request struct {
		Sku   string  `json:"sku"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	return func(ctxt *gin.Context) {
		var request request
		err := ctxt.BindJSON(&request)
		if err != nil {
			ctxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var p = domain.Product{
			Sku:   request.Sku,
			Name:  request.Name,
			Price: request.Price,
		}

		productId, err := c.productService.CreateNewProduct(p)
		if err, ok := err.(*pq.Error); ok {
			ctxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctxt.JSON(http.StatusOK, gin.H{"productId": productId})
	}
}

func (c *CatalogController) handleProductRead() gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		productId, err := uuid.Parse(ctxt.Param("id"))
		if err != nil {
			ctxt.JSON(http.StatusBadRequest, gin.H{"error": "Product ID must be a valid UUID"})
			return
		}

		p, err := c.productService.ReadProduct(productId)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				ctxt.JSON(http.StatusNotFound, gin.H{"error": "Could not find product"})
			default:
				ctxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		ctxt.JSON(http.StatusOK, gin.H{
			"data": p,
		})
	}
}
