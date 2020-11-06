package basic

import "testing"

// 合并数组测试
func TestMergeSlice(t *testing.T){
	var arr1 = []int{1,2,3}
	var arr2 = []int{4,5,6}
	var arr3 = []int{7,8,9}
	t.Logf("s1: %v\n", append(append(arr1, arr2...), arr3...))
}


