package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Sirupsen/logrus"
)

// TestMain - Main test function, tests start here
func TestMain(m *testing.M) {
	errCode := m.Run()
	err := os.Remove("tests/lastupdate.tmp")
	if err != nil {
		logrus.Error(err)
	}
	err = os.RemoveAll("tests/log")
	if err != nil {
		logrus.Error(err)
	}
	err = filepath.Walk("routers", func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), "commentsRouter_________________________________") {
			err = os.Remove("routers/" + info.Name())
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
	os.Exit(errCode)
}
