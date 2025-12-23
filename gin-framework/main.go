// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// NestJS: const app = express()
	// Gin: router := gin.Default()
	router := gin.Default() // With logger & recovery middleware

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Start server
	router.Run(":8080") // Listen on port 8080
}
