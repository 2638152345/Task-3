	products := r.Group("/api/v1/products")
	{
		products.GET("", listProducts)         
		products.GET("/:id", getProduct)       
		products.POST("", createProduct)     
		products.PUT("/:id", updateProduct)   
		products.DELETE("/:id", deleteProduct) 
	}

	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.PATCH("/users/:id", patchUser)
	r.DELETE("/users/:id", deleteUser)
	r.HEAD("/users", headUsers)
	r.OPTIONS("/users", optionsUsers)

	r.Any("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": c.Request.Method,
		})
	})

	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": []gin.H{
			{"id": 1, "name": "Alice"},
			{"id": 2, "name": "Bob"},
		},
	})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": "User " + id,
	})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User " + id + " updated",
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User " + id + " deleted",
	})
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"products": []gin.H{
			{"id": 1, "name": "Product 1"},
			{"id": 2, "name": "Product 2"},
		},
	})
}

func getProductsV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "v2",
		"products": []gin.H{
			{"id": 1, "name": "Product 1", "price": 99.99},
		},
	})
}

func listProducts(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"size":     size,
		"products": []gin.H{},
	})
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": "Product " + id,
	})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
	})
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product " + id + " updated",
	})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product " + id + " deleted",
	})
}

func patchUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User " + id + " patched",
	})
}

func headUsers(c *gin.Context) {
	c.Status(http.StatusOK)
}

func optionsUsers(c *gin.Context) {
	c.Header("Allow", "GET, POST, PUT, DELETE, OPTIONS")
	c.Status(http.StatusOK)
}
