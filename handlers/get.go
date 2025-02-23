package handlers

import (
	"net/http"
	"net/url"
	"sparrow/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get list of available security policies

// @Description Get the name of every loaded policy (securityPolicyId.name)
// @Success 200 {array} string "List of policy names" example:["policy1", "policy2"]
// @Router /api/policies [get]
func PoliciesHandler(spifs []utils.SPIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.GetPolicies(spifs))
	}
}

// @Summary Get list of available classifications
// @Description Get list of available classifications for a given policy (securityPolicyId.name)
// @Param policy path string true "Mandatory policy parameter"
// @Success 200 {array} string "List of classifications"
// @Failure 400  "Bad request"
// @Router /api/classifications/{policy} [get]
func ClassificationsHandler(spifs []utils.SPIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedPolicy, err := url.QueryUnescape(c.Param("policy"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, utils.GetClassifications(spifs, decodedPolicy))
	}
}

// @Summary Get list of security categories related to a policy and a classification
// @Description Get list of available security categories for a given policy (securityPolicyId.name)
// @Param policy path string true "Mandatory policy name"
// @Param classification path string true "Mandatory classification name"
// @Success 200 {array} string "List of security categories"
// @Failure 400  "Bad request"
// @Router /api/categories/{policy}/{classification} [get]
func CategoriesHandler(spifs []utils.SPIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedPolicy, errP := url.QueryUnescape(c.Param("policy"))
		decodedClassification, errC := url.QueryUnescape(c.Param("classification"))
		if errP != nil || errC != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, utils.GetCategories(spifs, decodedPolicy, decodedClassification))
	}
}

// @Summary Get the type of a security category
// @Description Get the type of a security category  (securityPolicyId.name) (securityPolicyId.name)
// @Param policy path string true "Mandatory policy parameter"
// @Param category path string true "Mandatory category parameter"
// @Success 200 {array} string "List of classifications"
// @Failure 400  "Bad request"
// @Router /api/type/{policy}/{category} [get]
func TypeHandler(spifs []utils.SPIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedPolicy, errP := url.QueryUnescape(c.Param("policy"))
		decodedCategory, errC := url.QueryUnescape(c.Param("category"))
		if errP != nil || errC != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, utils.GetType(spifs, decodedPolicy, decodedCategory))
	}
}
