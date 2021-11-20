package mathutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomInt(t *testing.T) {
	min, max := 100, 999

	for i := 0; i < 10; i++ {
		val := RandomInt(min, max)
		fmt.Println(val)

		assert.True(t, val >= min)
		assert.True(t, val <= max)
	}
}

func TestRealRandomInt(t *testing.T) {
	var min, max int64 = 100, 999

	for i := 0; i < 10; i++ {
		val := RealRandomInt(min, max)
		fmt.Println(val)

		assert.True(t, val >= min)
		assert.True(t, val <= max)
	}
}
