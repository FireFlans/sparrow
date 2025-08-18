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

// @Summary Returns a XML Label from a simplified JSON input
// @Description Returns a XML Label from a simplified JSON input
// @Success 200
// @Failure 400 Bad or missing label provided
// @Router /api/v1/generate [post]

func GenerateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonData structures.JSONConfidentialityLabel
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		xmlLabel := utils.GenerateXMLLabel(jsonData)

		c.Data(http.StatusOK, "application/xml", []byte(xmlLabel))
	}
}

// @Summary Returns a JSON dominant label from a list of JSON labels
// @Description Returns a JSON Dominant Label, computed from the rules described in SRD ADatP-4774.1, section 4.4
// @Success 200
// @Failure 400 Bad or missing label provided
// @Router /api/v1/dominant [post]

func DominantLabelHandler(spifs []structures.SPIF) gin.HandlerFunc {

	return func(c *gin.Context) {
		var listLabels []structures.JSONConfidentialityLabel
		if err := c.ShouldBindJSON(&listLabels); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(listLabels) < 2 {
			c.JSON(http.StatusBadRequest, "Error : Need to provide at least 2 labels")
		}
		label, err := utils.DominantLabel(spifs, listLabels)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, utils.ParseXMLLabel(label))

	}
}
