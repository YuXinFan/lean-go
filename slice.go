package helper

import (
  "fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func Append[T any](sl ...any) []T {
	var temp []T = nil
	var p T
	var s []T
	ptyp := reflect.TypeOf(p)
	styp := reflect.TypeOf(s)
	for _, v := range sl {
		if typ := reflect.TypeOf(v); typ == ptyp {
			temp = append(temp, v.(T))
		} else if typ == styp {
			temp = append(temp, v.([]T)...)
		} else {
			panic(fmt.Sprintf("#Lean# func Append[T any](sl ...any) []T # sl[i] type error %v", reflect.TypeOf(v)))
		}
	}
	return temp
}
