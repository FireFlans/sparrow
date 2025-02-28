package main

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestPolicies(t *testing.T) {
	Test(t,
		Description("Request Policies"),
		Get("http://localhost:8080/api/v1/policies"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("NATO"),
		Expect().Body().JSON().Contains("ACME"),
	)
}

func TestClassifications(t *testing.T) {
	Test(t,
		Description("Request ACME Classifications"),
		Get("http://localhost:8080/api/v1/classifications/ACME"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("PUBLIC"),
		Expect().Body().JSON().Contains("CONFIDENTIAL"),
		Expect().Body().JSON().Contains("INTERNAL"),
	)

	Test(t,
		Description("Request NATO Classifications"),
		Get("http://localhost:8080/api/v1/classifications/NATO"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("UNCLASSIFIED"),
		Expect().Body().JSON().Contains("RESTRICTED"),
		Expect().Body().JSON().Contains("CONFIDENTIAL"),
		Expect().Body().JSON().Contains("SECRET"),
		Expect().Body().JSON().Contains("TOP SECRET"),
	)

	Test(t,
		Description("Request Classifications non-existent policy"),
		Get("http://localhost:8080/api/v1/classifications/DUMMYACME"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal("null"),
	)
	Test(t,
		Description("Request Classifications without policy"),
		Get("http://localhost:8080/api/v1/classifications/"),
		Expect().Status().Equal(http.StatusNotFound),
	)
}

func TestCategories(t *testing.T) {
	Test(t,
		Description("Request categories for PUBLIC ACME"),
		Get("http://localhost:8080/api/v1/categories/ACME/PUBLIC"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().NotContains("Releasable To"),  // Should not appear
		Expect().Body().JSON().NotContains("Administrative"), // Should not appear
		Expect().Body().JSON().NotContains("Sensitive"),      // Should not appear
	)

	Test(t,
		Description("Request categories for CONFIDENTIAL ACME"),
		Get("http://localhost:8080/api/v1/categories/ACME/CONFIDENTIAL"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("Releasable To"),
		Expect().Body().JSON().NotContains("Administrative"), // Should not appear
		Expect().Body().JSON().NotContains("Sensitive"),      // Should not appear
	)

	Test(t,
		Description("Request categories for INTERNAL ACME"),
		Get("http://localhost:8080/api/v1/categories/ACME/INTERNAL"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().NotContains("Releasable To"), //Should not appear
		Expect().Body().JSON().Contains("Administrative"),
		Expect().Body().JSON().Contains("Sensitive"),
	)

	Test(t,
		Description("Request categories for UNCLASSIFIED NATO"),
		Get("http://localhost:8080/api/v1/categories/NATO/UNCLASSIFIED"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("Additional Sensitivity"),
		Expect().Body().JSON().Contains("Releasable To"),
		Expect().Body().JSON().Contains("Only"),
		Expect().Body().JSON().Contains("Administrative"),
		Expect().Body().JSON().Contains("Context"),
	)

	Test(t,
		Description("Request categories for TOP SECRET NATO"),
		Get("http://localhost:8080/api/v1/categories/NATO/TOP%20SECRET"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("Additional Sensitivity"),
		Expect().Body().JSON().Contains("Releasable To"),
		Expect().Body().JSON().Contains("Only"),
		Expect().Body().JSON().Contains("Administrative"),
		Expect().Body().JSON().Contains("Context"),
	)

	Test(t,
		Description("Request categories without policy"),
		Get("http://localhost:8080/api/v1/categories/UNRESTRCITED"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal("null"),
	)

	Test(t,
		Description("Request categories without classification"),
		Get("http://localhost:8080/api/v1/categories/NATO"),
		Expect().Status().Equal(http.StatusOK),
	)

	Test(t,
		Description("Request categories to non-existent policy"),
		Get("http://localhost:8080/api/v1/categories/NAO/UNCLASSIFIED"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal("null"),
	)

	Test(t,
		Description("Request categories to non-existent classification"),
		Get("http://localhost:8080/api/v1/categories/NATO/UNLASSIFIED"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("Additional Sensitivity"),
		Expect().Body().JSON().Contains("Releasable To"),
		Expect().Body().JSON().Contains("Only"),
		Expect().Body().JSON().Contains("Administrative"),
		Expect().Body().JSON().Contains("Context"),
	)

}
