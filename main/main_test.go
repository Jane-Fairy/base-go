package main

import (
	"fmt"
	"testing"
)

//func TestGetAare(t *testing.T) {
//	aare := GetAare(1, 2)
//	fmt.Println(aare)
//}

func BenchMarkGetAare(t *testing.B) {
	aare := GetAare(1, 2)
	fmt.Println(aare)
}
