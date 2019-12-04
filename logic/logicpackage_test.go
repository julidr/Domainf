package logic

import (
	"Domainf/models"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

// Test calculateLowerSslGrade function

func TestCalculateLowerSslGrade (test *testing.T) {
	testEndpoints := []models.Endpoint{{"1", "A"}, {"2", "F"},
		{"3", "B"}, {"4", "C"}}
	lowerGrade := calculateLowerSslGrade(testEndpoints)
	if lowerGrade != "F" {
		test.Errorf("Grade was incorrect, got: %v expected: %v", lowerGrade, "F")
	}
}

func TestCalculateGradeWithNotValidGrades (test *testing.T) {
	testEndpoints := []models.Endpoint{{"1", "R"}, {"2", "M"}}
	lowerGrade := calculateLowerSslGrade(testEndpoints)
	if lowerGrade != "" {
		test.Errorf("Both grades were not valid, it was expected an empty value not %v", lowerGrade)
	}
}

func TestCalculateGradeWithValidAndNotGrades (test *testing.T) {
	testEndpoints := []models.Endpoint{{"1", "A+"}, {"2", "M"}}
	lowerGrade := calculateLowerSslGrade(testEndpoints)
	if lowerGrade != "A+" {
		test.Errorf("Grade was incorrect, got: %v expected: %v", lowerGrade, "A+")
	}
}

// Test compareServers function

func TestCompareServerEquals(test *testing.T) {
	testRequestServers := []string{"1", "2", "3"}
	testDatabaseServers := []string{"1", "2", "3"}
	result := compareServers(testRequestServers, testDatabaseServers)
	if result != false {
		test.Errorf("Compare was wrong, both servers array are the same")
	}
}

func TestCompareServerDifferent(test *testing.T) {
	testRequestServers := []string{"1", "2", "3"}
	testDatabaseServers := []string{"1", "2", "4"}
	result := compareServers(testRequestServers, testDatabaseServers)
	if result != true {
		test.Errorf("Compare was wrong, servers were different")
	}
}

func TestCompareServerDifferentSize(test *testing.T) {
	testRequestServers := []string{"1", "2", "3"}
	testDatabaseServers := []string{"4", "5"}
	result := compareServers(testRequestServers, testDatabaseServers)
	if result != true {
		test.Errorf("Compare was wrong, servers were different in size")
	}
}

// Test searchIcon html

func TestSearchIconFound(test *testing.T) {
	htmlText := "<head><link rel=\"shortcut icon\" href=\"/favicon.ico\"></head>"
	htmlNode, _ := html.Parse(strings.NewReader(htmlText))
	icon := searchIcon(htmlNode)
	if icon != "/favicon.ico" {
		test.Errorf("Search icon failed, got: %v expected: %v", icon, "favicon.ico")
	}
}

func TestSearchIconNotFound(test *testing.T) {
	htmlText := "<head><link rel=\"icon\" href=\"/favicon.ico\"></head>"
	htmlNode, _ := html.Parse(strings.NewReader(htmlText))
	icon := searchIcon(htmlNode)
	if icon != "" {
		test.Errorf("No shorcut icon rel was specified the return should have been an empty string")
	}
}

// Test searchTitle html

func TestSearchTitleFound(test *testing.T) {
	htmlText := "<head><title>Test</title></head>"
	htmlNode, _ := html.Parse(strings.NewReader(htmlText))
	title := searchTitle(htmlNode)
	if title != "Test" {
		test.Errorf("Search title failed, got: %v expected: %v", title, "Title")
	}
}

func TestSearchTitleNotFound(test *testing.T) {
	htmlText := "<head></head>"
	htmlNode, _ := html.Parse(strings.NewReader(htmlText))
	title := searchIcon(htmlNode)
	if title != "" {
		test.Errorf("No title tag was specified the return should have been an empty string")
	}
}