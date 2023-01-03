package notes_loader

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestLoadNotes(t *testing.T) {
	setup := func() string {
		dir, err := os.MkdirTemp("", "terminote")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return dir
	}

	teardown := func(dir string) {
		// Remove the temporary directory
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	t.Run("empty directory", func(t *testing.T) {
		// Create a temporary directory
		dir := setup()
		defer teardown(dir)

		// Test loading notes from an empty directory
		expectedNotes := []Note{}
		actualNotes, err := LoadNotes(dir)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(actualNotes, expectedNotes) {
			t.Errorf("Unexpected notes: got %v, want %v", actualNotes, expectedNotes)
		}
	})

	t.Run("non-empty directory", func(t *testing.T) {
		// Create a temporary directory
		dir := setup()
		defer teardown(dir)

		// Create some files in the temporary directory
		files := []string{"note1.txt", "note2.txt"}
		for _, f := range files {
			_, err := os.Create(fmt.Sprintf("%s/%s", dir, f))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// Test loading notes from a non-empty directory
		expectedNotes := []Note{
			{Title: "note1.txt", Path: fmt.Sprintf("%s/note1.txt", dir)},
			{Title: "note2.txt", Path: fmt.Sprintf("%s/note2.txt", dir)},
		}
		actualNotes, err := LoadNotes(dir)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(actualNotes, expectedNotes) {
			t.Errorf("Unexpected notes: got %v, want %v", actualNotes, expectedNotes)
		}
	})

	t.Run("non-existent directory", func(t *testing.T) {
		notesDir := "testdata/nonexistent"
		_, err := LoadNotes(notesDir)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
