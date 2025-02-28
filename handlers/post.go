package handlers

import (
	"fmt"
	"net/http"
	"sparrow/structures"
	"sparrow/utils"

	"github.com/gin-gonic/gin"
)

func MarkingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		svgData := utils.GenerateSVG()
		c.Data(http.StatusOK, "image/svg+xml", svgData)
	}
}

func GenerateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		svgData := utils.GenerateSVG()
		c.Data(http.StatusOK, "image/svg+xml", svgData)
	}
}

// @Summary Returns a simplified JSON of the label
// @Description Returns a simplified JSON representation of the security label
// @Success 200
// @Failure 400 Bad or missing label provided
// @Router /api/v1/parse [post]

func ParseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var label structures.OriginatorConfidentialityLabel

		if err := c.ShouldBindXML(&label); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid XML format"})
			return
		}

		c.JSON(http.StatusOK, utils.ParseXMLLabel(label))
	}
}
