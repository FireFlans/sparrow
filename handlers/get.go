package handlers

import (
	"net/http"
	"net/url"
	"sparrow/structures"
	"sparrow/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Get list of available security policies
// @Description Get the name of every loaded policy (securityPolicyId.name)
// @Success 200 {array} string "List of policy names" example:["policy1", "policy2"]
// @Router /api/v1/policies [get]
func PoliciesHandler(spifs []structures.SPIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.GetPolicies(spifs))
	}
}

// @Summary Get list of available classifications
// @Description Get list of available classifications for a given policy (securityPolicyId.name)
// @Param policy path string true "Mandatory policy parameter"
// @Success 200 {array} string "List of classifications"
// @Failure 400  "Bad request"
// @Router /api/v1/classifications/{policy} [get]
func ClassificationsHandler(spifs []structures.SPIF) gin.HandlerFunc {
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
// @Param classification path string false "Optionnal classification name"
// @Success 200 {array} string "List of security categories"
// @Failure 400  "Bad request"
// @Router /api/v1/categories/{policy}/{classification} [get]
func CategoriesHandler(spifs []structures.SPIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedPolicy, errP := url.QueryUnescape(c.Param("policy"))
		decodedClassification, errC := url.QueryUnescape(c.Param("classification"))
		if decodedClassification != "" {
			decodedClassification = strings.TrimPrefix(decodedClassification, "/")
		}
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
// @Router /api/v1/type/{policy}/{category} [get]
func TypeHandler(spifs []structures.SPIF) gin.HandlerFunc {
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

// @Summary Get security mentions
// @Description Get the tagSecurity for a mention
// @Param policy path string true "Mandatory policy parameter"
// @Param classification path string true "Mandatory category parameter"
// @Param category path string true "Mandatory category parameter"
// @Success 200 {array} string "List of classifications"
// @Failure 400  "Bad request"
// @Router /api/v1/mentions/{policy}/{classification}/{category} [get]
func MentionsHandler(spifs []structures.SPIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedPolicy, errP := url.QueryUnescape(c.Param("policy"))
		decodedCategory, errCa := url.QueryUnescape(c.Param("category"))
		decodedClassification, errCl := url.QueryUnescape(c.Param("classification"))
		if errP != nil || errCa != nil || errCl != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, utils.GetMentions(spifs, decodedPolicy, decodedClassification, decodedCategory))
	}
}
