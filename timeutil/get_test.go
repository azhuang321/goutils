package timeutil

import (
	"fmt"
	"testing"
)

func TestGetNow(t *testing.T) {
	fmt.Println(GetNow())
	fmt.Println(GetNowTimestamp())
}
