package utilities

import (
	"os"
	"testing"
)

func CreateTestFile(t *testing.T, data string) (*os.File, func()) {
	tempFile, err := os.CreateTemp("", "outputs.json")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}

	if data != "" {
		_, err = tempFile.WriteString(data)
		if err != nil {
			t.Fatalf("Error writing test data to file: %v", err)
		}
	}

	cleanup := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}
	return tempFile, cleanup
}
