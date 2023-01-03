package notes_loader

import (
	"fmt"
	"os"
)

type Note struct {
	Title string
	Path  string
}

func LoadNotes(notesDir string) ([]Note, error) {
	dir, err := os.Open(notesDir)
	if err != nil {
		return []Note{}, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return []Note{}, err
	}

	notes := []Note{}
	for _, f := range files {
		notes = append(notes, Note{
			Title: f.Name(),
			Path:  fmt.Sprintf("%s/%s", notesDir, f.Name()),
		})
	}
	return notes, nil
}
