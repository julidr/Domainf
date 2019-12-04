package utils

import "testing"

// Test FormatServersToDBArray function

func TestFormatServersToDBArray(test *testing.T) {
	testServers := []string{"1", "2", "3"}
	result := FormatServersToDBArray(testServers)
	expectedResult := "ARRAY['1','2','3']"
	if result != expectedResult {
		test.Errorf("FormatServersToDBArray failed, got: %v expected: %v", result, expectedResult)
	}
}

func TestFormatServersToDBArrayEmpty(test *testing.T) {
	testServers := []string{}
	result := FormatServersToDBArray(testServers)
	expectedResult := "ARRAY[]"
	if result != expectedResult {
		test.Errorf("FormatServersToDBArray failed, got: %v expected: %v", result, expectedResult)
	}
}

// Test FormatDBArrayToServers function

func TestFormatDBArrayToServers(test *testing.T) {
	testArray := "{1,2,3}"
	result := FormatDBArrayToServers(testArray)
	expectedResult := []string{"1", "2", "3"}
	for i := 0; i < len(result); i++ {
		if result[i] != expectedResult[i] {
			test.Errorf("FormatDBArrayToServers failed, got: %v expected: %v", result, expectedResult)
		}
	}
}