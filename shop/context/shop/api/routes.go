package api

func (c *ShopController) Run() {
	// Customer
	c.router.GET("/customers/:id", c.handleCustomerRead())
	c.router.POST("/customers", c.handleCustomerCreate())

	// Order
	c.router.POST("/orders", c.handleOrderCreate())

	// Sales
	c.router.GET("/sales", c.handleSalesRead())
}
