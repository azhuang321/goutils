package arrutil

import (
	"errors"
	"goutils/mathutil"
	"goutils/strutil"
	"reflect"
	"strconv"
)

var ErrInvalidType = errors.New("the input param type is invalid")

// ToInt64s 转换为 []int64
// @param arr
// @return ret
// @return err
func ToInt64s(arr interface{}) (ret []int64, err error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		err = ErrInvalidType
		return
	}

	for i := 0; i < rv.Len(); i++ {
		i64, err := mathutil.Int64(rv.Index(i).Interface())
		if err != nil {
			return []int64{}, err
		}
		ret = append(ret, i64)
	}
	return
}

// MustToInt64s 一定可以转换为 []int64
// @param arr
// @return []int64
func MustToInt64s(arr interface{}) []int64 {
	ret, _ := ToInt64s(arr)
	return ret
}

// SliceToInt64s slice interface 转换为 []int64
// @param arr
// @return []int64
func SliceToInt64s(arr []interface{}) []int64 {
	i64s := make([]int64, len(arr))
	for i, v := range arr {
		i64s[i] = mathutil.MustInt64(v)
	}
	return i64s
}

// ToStrings 转换 interface{}(allow: array,slice) 为 []string
// @param arr
// @return ret
// @return err
func ToStrings(arr interface{}) (ret []string, err error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		err = ErrInvalidType
		return
	}

	for i := 0; i < rv.Len(); i++ {
		str, err := strutil.ToString(rv.Index(i).Interface())
		if err != nil {
			return []string{}, err
		}

		ret = append(ret, str)
	}
	return
}

// MustToStrings 一定可以转换 interface{}(allow: array,slice) 为 []string
// @param arr
// @return []string
func MustToStrings(arr interface{}) []string {
	ret, _ := ToStrings(arr)
	return ret
}

// SliceToStrings 转换 []interface{} 为 []string
// @param arr
// @return []string
func SliceToStrings(arr []interface{}) []string {
	ss := make([]string, len(arr))
	for i, v := range arr {
		ss[i] = strutil.MustString(v)
	}
	return ss
}

// StringsToInts []string 转换为 []int
// @param ss
// @return ints
// @return err
func StringsToInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}
