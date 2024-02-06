package main

import (
	"reflect"

	"github.com/wxnacy/wgo/arrays"

	"slices"
)

func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

func StringsContains(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func SliceContains(array []string, val string) (index int) {
	i := slices.Index(array, val) // Use .Index() here, because other .Contains() methods return index, not boolean
	return i
}

func WgoContains(array []string, val string) (index int) {
	i := arrays.Contains(array, val)
	return i
}

func WgoSwitchContains(array interface{}, val interface{}) (index int) {
	var i int
	switch val.(type) {
	case string:
		var arrtmp []string
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			arrtmp = append(arrtmp, s.Index(i).Interface().(string))
		}
		i = arrays.StringsContains(arrtmp, val.(string))
	}

	return i
}

func WgoStringsContains(array []string, val string) (index int) {
	i := arrays.StringsContains(array, val)
	return i
}

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

// func StringsContains(array []string, val string) (index int) {
// 	index = -1
// 	for i := 0; i < len(array); i++ {
// 		if array[i] == val {
// 			index = i
// 			return
// 		}
// 	}
// 	return
// }

// func BenchmarkContains(b *testing.B) {
// 	sa := []string{"q", "w", "e", "r", "t"}
// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		Contains(sa, "r")
// 	}
// }

// func BenchmarkStringsContains(b *testing.B) {
// 	sa := []string{"q", "w", "e", "r", "t"}
// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		StringsContains(sa, "r")
// 	}
// }
