package main

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	invalidURL = "http://this-is-an-invalid-url"
	validURL   = "https://en.wikipedia.org/wiki/Robotics"
)

func TestPullURL(t *testing.T) {
	assert := assert.New(t)

	// Test pullURL with invalid URL
	_, _, _, err := pullURL(invalidURL)
	assert.Error(err, "pullURL should return an error for invalid URLs")

	// Test pullURL with valid URL
	_, title, text, err := pullURL(validURL)
	assert.NoError(err, "pullURL should not return an error for valid URLs")
	assert.NotEmpty(title, "Title should not be empty")
	assert.NotEmpty(text, "Text should not be empty")
}

func TestWriteToHTML(t *testing.T) {
	assert := assert.New(t)

	// Test writeToHTML with invalid file path
	err := writeToHTML(":", []byte("Test HTML Data"))
	assert.Error(err, "writeToHTML should return an error for invalid file paths")

	// Test writeToHTML with valid data
	err = writeToHTML(validURL, []byte("Test HTML Data"))
	assert.NoError(err, "writeToHTML should not return an error for valid data")

	// Read the written file and check its contents
	writtenData, err := os.ReadFile("wikipages/" + path.Base(validURL) + ".html")
	assert.NoError(err, "Reading the written file should not return an error")
	assert.Equal(string(writtenData), "Test HTML Data", "The written data does not match")

	// Clean up the created file
	err = os.Remove("wikipages/" + path.Base(validURL) + ".html")
	assert.NoError(err, "Removing the written file should not return an error")
}

func TestWriteToJL(t *testing.T) {
	assert := assert.New(t)

	// Backup existing goItems.jl file
	err := os.Rename("goItems.jl", "goItems.jl.bak")
	if err != nil && !os.IsNotExist(err) {
		assert.FailNow("Failed to backup goItems.jl", err.Error())
	}
	defer func() {
		// Restore the backup file
		err := os.Rename("goItems.jl.bak", "goItems.jl")
		if err != nil && !os.IsNotExist(err) {
			assert.FailNow("Failed to restore goItems.jl", err.Error())
		}
	}()

	// Test writeToJL with invalid file path
	err = writeToJL(":", "Test Title", "Test Text")
	assert.Error(err, "writeToJL should return an error for invalid file paths")

	// Test writeToJL with valid data
	err = writeToJL(validURL, "Test Title", "Test Text")
	assert.NoError(err, "writeToJL should not return an error for valid data")

	// Read the written file and check its contents
	writtenData, err := os.ReadFile("goItems.jl")
	assert.NoError(err, "Reading the written file should not return an error")
	expectedData := `{"url":"` + validURL + `","title":"Test Title","text":"Test Text"}` + "\n"
	assert.True(strings.Contains(string(writtenData), expectedData), "The written data does not match")

	// Clean up the created file
	err = os.Remove("goItems.jl")
	assert.NoError(err, "Removing the written file should not return an error")
}
