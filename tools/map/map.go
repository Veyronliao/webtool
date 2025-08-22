package _map

import (
	"cmp"
	stdmaps "maps"
	"slices"
)

func GetMapAllKeys[K comparable, V any, M ~map[K]V](m M) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func SortedKeys[K cmp.Ordered, V any, M ~map[K]V](m M) []K {
	ks := GetMapAllKeys(m)
	slices.Sort(ks)
	return ks
}

func GetMapAllVals[K comparable, V any, M ~map[K]V](m M) []V {
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func SortedMapVals[K comparable, V cmp.Ordered, M ~map[K]V](m M) []V {
	s := GetMapAllVals(m)
	slices.Sort(s)
	return s
}

// 以两个数组返回map的keys和values
func GetMapKeysAndVals[K comparable, V any, M ~map[K]V](m M) ([]K, []V) {
	ks, vs := make([]K, 0, len(m)), make([]V, 0, len(m))
	for k, v := range m {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return ks, vs
}

// 合并n个map
func MergeMaps[K comparable, V any, M ~map[K]V](ms ...M) map[K]V {
	if len(ms) == 0 {
		return nil
	}

	m := make(map[K]V)
	for _, v := range ms {
		stdmaps.Copy(m, v)
	}
	return m
}

func IsContainsAllKeys[K comparable, V any, M ~map[K]V](m M, keys []K) bool {
	for _, key := range keys {
		if _, exists := m[key]; !exists {
			return false
		}
	}
	return true
}

// 判断map中是否存在任意的key
func IsContainsAnyKey[K comparable, V any, M ~map[K]V](m M, keys []K) bool {
	for _, key := range keys {
		if _, exists := m[key]; exists {
			return true
		}
	}
	return false
}
