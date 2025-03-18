package main

import (
	"go/web/controllers/c_internal"
	"go/web/middlewares"

	// "./controllers/internal"
	"github.com/gin-gonic/gin"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

func main() {
	router := gin.Default()
	router.Use(middlewares.CorsHeader())

	router.POST("/internal/login", c_internal.Login)

	// r_internal := router.Group("/internal")
	// r_internal.GET("/login", c_internal.Login)

	// r_main := router.Group("/main")
	// r_main.GET("/login", c_main.Login)

	// router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
