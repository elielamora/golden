package golden

import (
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = os.Getenv("UPDATE_GOLDEN") == "true"
var testdataDir = "testdata" // always store golden files in the conventional directory
var goldenFilePermissions os.FileMode = 0644
var directoryPermissions os.FileMode = 0750

// ForceUpdate override the golden file update behavior
func ForceUpdate() {
	update = true
}

// Assert compares the golden file to the actual data.
// If the golden files are being updated then it writes and always passes.
func Assert(t *testing.T, actual []byte, msgAndArgs ...any) {
	t.Helper()
	expected := Value(t, actual)
	assert.Equal(t, expected, actual, msgAndArgs...)
}

// Value gets the raw data from the filesystem
func Value(t *testing.T, actual []byte) []byte {
	t.Helper()
	name := t.Name()
	dir, filename, _ := CutRight(name, "/")
	dir = path.Join(testdataDir, dir)
	if update {
		// ensure directory exists when updating
		err := os.MkdirAll(dir, directoryPermissions)
		require.NoError(t, err, "unexpected error creating directory %s: %s", dir, err)
	}
	// check if file has a preferred extension e.g. json/sql for syntax highlighting if not default to .golden
	if !strings.Contains(filename, ".") {
		filename += ".golden"
	}
	path.Join(testdataDir, dir, filename)
	goldenPath := path.Join(dir, filename)

	flags := os.O_RDONLY
	if update {
		flags = os.O_RDWR
		flags |= os.O_CREATE
		flags |= os.O_TRUNC // make sure we truncate the file since we'll always be rewriting the entire file
	}

	f, err := os.OpenFile(goldenPath, flags, goldenFilePermissions)
	if err != nil {
		require.NoError(t, err, "unexpected error opening file %s: %s", goldenPath, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			require.NoError(t, err, "unexpected error closing file %s: %s", goldenPath, err)
		}
	}()

	if update {
		_, err := f.Write(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	return content
}

// CutRight like strings.Cut but from the end.
func CutRight(s, sep string) (before, after string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return "", s, false
}
