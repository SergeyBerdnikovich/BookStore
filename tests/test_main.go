package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	err := m.Run()
	os.Remove("tests/lastupdate.tmp")
	os.RemoveAll("tests/log")
	filepath.Walk("routers", func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), "commentsRouter_________________________________") {
			os.Remove("routers/" + info.Name())
		}
		return nil
	})
	os.Exit(err)
}
