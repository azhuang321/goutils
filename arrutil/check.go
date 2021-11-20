package arrutil

import (
	"goutils/mathutil"
	"reflect"
	"strings"
)

// IntsHas int切片中是否包含给定值
// @param ints
// @param val
// @return bool
func IntsHas(intS []int, val int) bool {
	for _, ele := range intS {
		if ele == val {
			return true
		}
	}
	return false
}

// Int64sHas int切片中是否包含给定值
// @param ints
// @param val
// @return bool
func Int64sHas(ints []int64, val int64) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// StringsHas 字符串中是否包含给定值
// @param ss
// @param val
// @return bool
func StringsHas(ss []string, val string) bool {
	for _, ele := range ss {
		if ele == val {
			return true
		}
	}
	return false
}

// Contains interface的包含
// @param arr
// @param val
// @return bool
func Contains(arr, val interface{}) bool {
	if val == nil || arr == nil {
		return false
	}
	if strVal, ok := val.(string); ok {
		if ss, ok := arr.([]string); ok {
			return StringsHas(ss, strVal)
		}
		rv := reflect.ValueOf(arr)
		if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
			for i := 0; i < rv.Len(); i++ {
				if v, ok := rv.Index(i).Interface().(string); ok && strings.EqualFold(v, strVal) {
					return true
				}
			}
		}
		return false
	}

	intVal, err := mathutil.Int64(val)
	if err != nil {
		return false
	}
	if int64s, err := ToInt64s(arr); err == nil {
		return Int64sHas(int64s, intVal)
	}
	return false
}

// HasValue Contains的别名
// @param arr
// @param val
// @return bool
func HasValue(arr, val interface{}) bool {
	return Contains(arr, val)
}

// NotContains 不包含 某个值
// @param arr
// @param val
// @return bool
func NotContains(arr, val interface{}) bool {
	return false == Contains(arr, val)
}
