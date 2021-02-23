package basic

import (
	"testing"
)

func TestDelete(t *testing.T) {
	
	array :=[...]string{"a","b"}
	t.Log(&array)
	array[0] = "abc"
	
	
	var arrayString []string
	t.Logf("%p  \n" , arrayString)
	
	arrayString = append(arrayString, "qwer")
	t.Logf("%p %v \n" , arrayString,arrayString)
	
	
	
	
	// 使用 append  会改变原有数组的指针地址
	arrayString = append(arrayString, "asdfghj")
	t.Logf("%p  %v \n" , &arrayString,arrayString)
	
}
