package mathutil

import (
	"math/big"
	"time"

	cRand "crypto/rand"
	mRand "math/rand"
)

// RandomInt 生成基于微秒时间的伪随机数
// Usage:
// 	RandomInt(10, 99)
// 	RandomInt(100, 999)
// 	RandomInt(1000, 9999)
func RandomInt(min, max int) int {
	mRand.Seed(time.Now().UnixNano())
	return min + mRand.Intn(max-min)
}

// RealRandomInt 生成真随机数,但是性能 相较于 math/rand 差上10倍
// @reference https://zhuanlan.zhihu.com/p/94684495
// @param min
// @param max
// @return int64
func RealRandomInt(min, max int64) int64 {
	n, _ := cRand.Int(cRand.Reader, big.NewInt(max-min))
	return min + n.Int64()
}
