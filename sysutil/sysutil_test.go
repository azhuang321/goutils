package sysutil_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"goutils/sysutil"
)

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, sysutil.ProcessExists(pid))
}
