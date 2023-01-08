package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmdavril/template/newshop/app"
	"github.com/jmdavril/template/newshop/context/shop/domain"
	"github.com/lib/pq"
	"net/http"
)

type ShopController struct {
	router          *gin.Engine
	customerService *domain.CustomerService
}

func NewShopController(r *gin.Engine, s *domain.CustomerService) *ShopController {
	return &ShopController{
		router:          r,
		customerService: s,
	}
}

func (c *ShopController) handleCustomerCreate() gin.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}

	return func(ctxt *gin.Context) {
		var request request
		err := ctxt.BindJSON(&request)
		if err != nil {
			ctxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cusId, err := c.customerService.CreateNewCustomer(domain.Customer{Email: request.Email})
		if err, ok := err.(*pq.Error); ok {
			if err.Code == app.UniqueConstraintViolated {
				ctxt.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("customer with email '%v' already exists", request.Email)})
				return
			}
		}
		if err != nil {
			ctxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctxt.JSON(http.StatusOK, gin.H{"customerId": cusId})
	}
}

func (c *ShopController) handleCustomerRead() gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		customerId, err := uuid.Parse(ctxt.Param("id"))
		if err != nil {
			ctxt.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID must be a valid UUID"})
			return
		}

		cus, err := c.customerService.ReadCustomer(customerId)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				ctxt.JSON(http.StatusNotFound, gin.H{"error": "Could not find customer"})
			default:
				ctxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		ctxt.JSON(http.StatusOK, gin.H{
			"data": cus,
		})
	}
}

func (c *ShopController) handleOrderCreate() gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		var request OrderRequest
		err := ctxt.BindJSON(&request)
		if err != nil {
			ctxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		o, err := request.order()
		if err != nil {
			ctxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderId, err := c.customerService.CreateNewOrder(o)

		if err != nil {
			ctxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctxt.JSON(http.StatusOK, gin.H{"orderId": orderId})
	}
}

func (c *ShopController) handleSalesRead() gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		ctxt.JSON(http.StatusOK, gin.H{
			"data": "todo",
		})
	}
}
