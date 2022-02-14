package tests

import (
	"ebooks/pkg/models"
	"os"
	"testing"
)

var testAuthor models.Author

func TestMain(m *testing.M) {
	testAuthor = models.Author{
		FirstName:  "Zeus",
		MiddleName: "God of the",
		LastName:   "Sky",
	}
	testRunner := m.Run()
	os.Exit(testRunner)
}

func TestAuthorStructHasFirstName(t *testing.T) {
	expected := "Zeus"
	actual := testAuthor.FirstName
	if expected != actual {
		t.Errorf("Test failed: expected value: %q, actual value: %q", expected, actual)
	}
}

func TestAuthorStructHasLastName(t *testing.T) {
	expected := "Sky"
	actual := testAuthor.LastName
	if expected != actual {
		t.Errorf("Test failed: expected value: %q, actual value: %q", expected, actual)
	}
}
