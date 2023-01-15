package archive

import (
	"os"
	"path/filepath"
	"testing"
)

func TestArchiveDir(t *testing.T) {
	// Create sample files and dirs to archive
	if err := os.MkdirAll("test_data", 0755); err != nil {
		t.Errorf(err.Error())
	}
	dirs := []string{"dir1", "dir2"}
	for _, dirname := range dirs {
		if err := os.MkdirAll(filepath.Join("test_data", dirname), 0755); err != nil {
			t.Errorf(err.Error())
		}
	}
	files := []string{"file1", "file2", "dir1/file1", "dir1/file2", "dir2/file1", "dir2/file2"}
	for _, filename := range files {
		f, err := os.Create(filepath.Join("test_data", filename))
		if err != nil {
			t.Errorf(err.Error())
		}
		// write filename to file, check later if unpacked file content is still intact
		if _, err := f.WriteString(filename); err != nil {
			t.Errorf(err.Error())
		}
		if err := f.Close(); err != nil {
			t.Errorf(err.Error())
		}
	}

	if err := Archive("test_data/"); err != nil {
		t.Errorf(err.Error())
	}
	
	if err := Unarchive("test_data.tar.gz", "unpacked_test_data"); err != nil {
		t.Errorf(err.Error())
	}
	
	// TODO, check if unarchived files still has the same files
}
