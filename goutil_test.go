package goutils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"goutils"
)

func TestFuncName(t *testing.T) {
	name := goutils.FuncName(goutils.PkgName)
	assert.Equal(t, "goutils.PkgName", name)

	name = goutils.FuncName(goutils.PanicIfErr)
	assert.Equal(t, "goutils.PanicIfErr", name)
}

func TestPkgName(t *testing.T) {
	name := goutils.PkgName(goutils.FuncName(goutils.PanicIfErr))
	assert.Equal(t, "goutils", name)
}

func TestPanicIfErr(t *testing.T) {
	goutils.PanicIfErr(nil)
}

func TestPanicf(t *testing.T) {
	assert.Panics(t, func() {
		goutils.Panicf("hi %s", "inhere")
	})
}

func TestGetCallStacks(t *testing.T) {
	msg := goutils.GetCallStacks(false)
	fmt.Println(string(msg))

	fmt.Println("-------------full stacks-------------")
	msg = goutils.GetCallStacks(true)
	fmt.Println(string(msg))
}
