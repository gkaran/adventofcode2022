package utils

import (
	"golang.org/x/exp/constraints"
)

func Intersection[T constraints.Ordered](pS ...[]T) []T {
	hash := make(map[T]*int) // value, counter
	result := make([]T, 0)
	for _, slice := range pS {
		duplicationHash := make(map[T]bool) // duplication checking for individual slice
		for _, value := range slice {
			if _, isDup := duplicationHash[value]; !isDup { // is not duplicated in slice
				if counter := hash[value]; counter != nil { // is found in hash counter map
					if *counter++; *counter >= len(pS) { // is found in every slice
						result = append(result, value)
					}
				} else { // not found in hash counter map
					i := 1
					hash[value] = &i
				}
				duplicationHash[value] = true
			}
		}
	}
	return result
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

//func Intersection[T constraints.Ordered](a, b []T) (res []T) {
//	m := make(map[T]bool)
//
//	for _, item := range a {
//		m[item] = true
//	}
//
//	for _, item := range b {
//		if _, ok := m[item]; ok {
//			res = append(res, item)
//		}
//	}
//
//	slices.Sort(res)
//	return slices.Compact(res)
//}
