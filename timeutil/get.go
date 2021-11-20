//reference: https://github.com/golang-module/carbon

package timeutil

import (
	"fmt"

	"github.com/golang-module/carbon/v2"
)

func GetNow() string {
	return fmt.Sprintf("%s", carbon.Now())
}

func GetNowTimestamp() int64 {
	return carbon.Now().Timestamp()
}
