package handlers

import (
	"net/http"
	"sparrow/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Return marking from security label
// @Description Return marking from security label
// @Param type path string true "Desired type" Enums(png, svg)
// @Success 200
// @Router /api/v1/marking/{type} [post]
func MarkingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		svgData := utils.GenerateSVG()
		c.Data(http.StatusOK, "image/svg+xml", svgData)
	}
}
