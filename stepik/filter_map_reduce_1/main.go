package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := Filter(s, func(i int) bool {
		if i%2 == 0 {
			return true
		}
		return false
	},
	)
	fmt.Println(evens)
	//----------------
	sStr := Map(s, func(a int) string {
		return "'" + strconv.Itoa(a) + "'"
	})
	fmt.Println(sStr)
	sum := Reduce(s, 0, func(a int, b int) int {
		return a + b
	},
	)
	fmt.Println(sum)
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {

		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v1 := range s {
		r[i] = f(v1)

	}
	return r
}

func Reduce[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	r := init
	for _, v := range s {
		r = f(v, r)
	}
	return r
}
