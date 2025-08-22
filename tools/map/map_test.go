package _map

import (
	"fmt"
	"testing"
)

func TestGetMapAllKeys(t *testing.T) {
	ks := GetMapAllKeys(map[int]int{1: 1, 2: 2, 3: 3})
	if len(ks) == 3 {
		fmt.Printf("ks: %v\n", ks)
		fmt.Printf("pass")
	}
}
func TestSortedVals(t *testing.T) {

}
func TestSortedKeys(t *testing.T) {

}
func TestGetMapAllVals(t *testing.T) {

}
func TestSortedMapVals(t *testing.T) {

}
func TestSortedMapKeys(t *testing.T) {}

func TestGetMapKeysAndVals(t *testing.T) {

}

func TestMergeMaps(t *testing.T) {

}

func TestIsContainsAllKeys(t *testing.T) {

}
func TestIsContainsAnyKeys(t *testing.T) {

}
