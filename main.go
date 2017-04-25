package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	r "javierlvelasquez.com/resistor-color-code-lib/resistance"
	"net/http"
)

type Bands struct {
	Bands []string `json:"bands" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/tolerance", func(c *gin.Context) {
		var bands Bands
		if c.BindJSON(&bands) == nil {
			c.JSON(http.StatusOK, r.GetResistance(bands.Bands))
		}
	})
	router.Run(":8080")
}
