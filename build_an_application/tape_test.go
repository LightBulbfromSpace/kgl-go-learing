package poker_test

import (
	poker "github.com/LightBulbfromSpace/build_an_application"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "123456")

	defer clean()

	tmp := &poker.Tape{File: file}
	tmp.Write([]byte("smth"))

	file.Seek(0, 0)

	newFileContent, _ := io.ReadAll(file)
	want := "smth"
	assert.Equal(t, want, string(newFileContent))
}
