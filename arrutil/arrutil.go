package arrutil

import (
	"goutils/mathutil"
	"reflect"
	"strings"
)

// Reverse 反转字符串slice
// @param ss 反转的字符切片
func Reverse(ss []string) {
	l := len(ss)

	for i := 0; i < l/2; i++ {
		li := l - i - 1
		ss[i], ss[li] = ss[li], ss[i]
	}
}

// StringsRemove 移除字符切片中的特定字符
// @param ss 目标切片
// @param s 特定字符串
// @return []string 返回新的字符切片
func StringsRemove(ss []string, s string) []string {
	var ns []string
	for _, v := range ss {
		if v != s {
			ns = append(ns, v)
		}
	}
	return ns
}

// TrimStrings 去除多个字符串
// @param ss
// @param cutSet
// @return ns
func TrimStrings(ss []string, cutSet ...string) (ns []string) {
	hasCutSet := len(cutSet) > 0 && cutSet[0] != ""

	for _, str := range ss {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, cutSet[0]))
		} else {
			ns = append(ns, strings.TrimSpace(str))
		}
	}
	return
}

// GetRandomOne 从 切片中随机获取一个元素
// @param arr
// @return interface{}
func GetRandomOne(arr interface{}) interface{} {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return arr
	}

	i := mathutil.RandomInt(0, rv.Len())
	r := rv.Index(i).Interface()

	return r
}
