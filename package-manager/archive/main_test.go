package archive

import "testing"

func TestArchiveDir(t *testing.T) {
	if err := ArchiveDir("test_data/"); err != nil {
		t.Errorf(err.Error())
	}
}
