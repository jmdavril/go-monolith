package api

func (c *CatalogController) Run() {
	c.router.GET("/products/:id", c.handleProductRead())
	c.router.POST("/products", c.handleProductCreate())
}
