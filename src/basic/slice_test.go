package basic

import (
	"fmt"
	"testing"
)

// 合并数组测试
func TestMergeSlice(t *testing.T){
	var arr1 = []int{1,2,3}
	var arr2 = []int{4,5,6}
	var arr3 = []int{7,8,9}
	t.Logf("s1: %v\n", append(append(arr1, arr2...), arr3...))
}


func TestMains(t *testing.T) {
	array := [5]int{1: 2, 3:4}
	modify(array)
	fmt.Println(array)
}
// 可以得出，在函数传递时，数组传递的是 副本，所以函数修改副本内容时，原值不会收到影响
func modify(a [5]int){
	a[1] =3
	fmt.Println(a)
}
