package main

import (
	"encoding/json"
	"net/http"
	"os"
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

func TestMentions(t *testing.T) {
	Test(t,
		Description("Request mentions for PUBLIC ACME ADMINISTRATIVE"),
		Get("http://localhost:8080/api/v1/mentions/ACME/PUBLIC/Administrative"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal("null"),
		Expect().Body().JSON().NotContains("FINANCE"),     // Should not appear
		Expect().Body().JSON().NotContains("SALES"),       // Should not appear
		Expect().Body().JSON().NotContains("ENGINEERING"), // Should not appear
	)
	Test(t,
		Description("Request mentions for INTERNAL ACME ADMINISTRATIVE"),
		Get("http://localhost:8080/api/v1/mentions/ACME/INTERNAL/Administrative"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("FINANCE"),
		Expect().Body().JSON().Contains("SALES"),
		Expect().Body().JSON().Contains("ENGINEERING"),
	)
	Test(t,
		Description("Request mentions and forget a parameter"),
		Get("http://localhost:8080/api/v1/mentions/INTERNAL/Releasable%20To"),
		Expect().Status().Equal(http.StatusNotFound),
	)
	Test(t,
		Description("Request mentions for a non existent policy"),
		Get("http://localhost:8080/api/v1/mentions/ACE/PUBLIC/Administrative"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal("null"),
		Expect().Body().JSON().NotContains("FINANCE"),     // Should not appear
		Expect().Body().JSON().NotContains("SALES"),       // Should not appear
		Expect().Body().JSON().NotContains("ENGINEERING"), // Should not appear
	)

	Test(t,
		Description("Request mentions for a non existent classification"),
		Get("http://localhost:8080/api/v1/mentions/ACME/PULIC/Administrative"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Contains("FINANCE"),
		Expect().Body().JSON().Contains("SALES"),
		Expect().Body().JSON().Contains("ENGINEERING"),
	)

	Test(t,
		Description("Request mentions for a non existent mention"),
		Get("http://localhost:8080/api/v1/mentions/ACME/PUBLIC/Adminstrative"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal("null"),
		Expect().Body().JSON().NotContains("FINANCE"),     // Should not appear
		Expect().Body().JSON().NotContains("SALES"),       // Should not appear
		Expect().Body().JSON().NotContains("ENGINEERING"), // Should not appear
	)
}

func TestParsing(t *testing.T) {
	xmlData1, err := os.ReadFile("labels/label1.xml")
	if err != nil {
		t.Fatalf("Failed to read label1.xml: %v", err)
	}
	jsonData1, _ := os.ReadFile("labels/label1_simplified.json")
	var expectedJSONResponse1 map[string]interface{}
	err = json.Unmarshal(jsonData1, &expectedJSONResponse1)
	if err != nil {
		t.Fatalf("Failed to read labels/label1_simplified.json: %v", err)
	}
	Test(t,
		Description("Request JSON representation for label1"),
		Post("http://localhost:8080/api/v1/parse"),
		Send().Headers("Content-Type").Add("application/xml"),
		Send().Body().String(string(xmlData1)),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(expectedJSONResponse1),
	)

	xmlData2, err := os.ReadFile("labels/label2.xml")
	if err != nil {
		t.Fatalf("Failed to read label2.xml: %v", err)
	}
	jsonData2, _ := os.ReadFile("labels/label2_simplified.json")
	var expectedJSONResponse2 map[string]interface{}
	err = json.Unmarshal(jsonData2, &expectedJSONResponse2)
	if err != nil {
		t.Fatalf("Failed to read labels/label2_simplified.json: %v", err)
	}
	Test(t,
		Description("Request JSON representation for label2"),
		Post("http://localhost:8080/api/v1/parse"),
		Send().Headers("Content-Type").Add("application/xml"),
		Send().Body().String(string(xmlData2)),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(expectedJSONResponse2),
	)
}

func TestGenerating(t *testing.T) {
	xmlData1, err := os.ReadFile("labels/label1.xml")
	if err != nil {
		t.Fatalf("Failed to read label1.xml: %v", err)
	}
	jsonData1, _ := os.ReadFile("labels/label1_simplified.json")
	var jsonBody1 map[string]interface{}
	err = json.Unmarshal(jsonData1, &jsonBody1)
	if err != nil {
		t.Fatalf("Failed to read labels/label2_simplified.json: %v", err)
	}
	Test(t,
		Description("Request XML representation for label1"),
		Post("http://localhost:8080/api/v1/generate"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(jsonBody1),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal(string(xmlData1)),
	)

	xmlData2, err := os.ReadFile("labels/label2.xml")
	if err != nil {
		t.Fatalf("Failed to read label2.xml: %v", err)
	}
	jsonData2, _ := os.ReadFile("labels/label2_simplified.json")
	var jsonBody2 map[string]interface{}
	err = json.Unmarshal(jsonData2, &jsonBody2)
	if err != nil {
		t.Fatalf("Failed to read labels/label2_simplified.json: %v", err)
	}
	Test(t,
		Description("Request XML representation for label2"),
		Post("http://localhost:8080/api/v1/generate"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(jsonBody2),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(string(xmlData2)),
	)
}
