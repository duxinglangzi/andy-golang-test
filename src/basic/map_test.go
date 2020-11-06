package basic

import (
	"encoding/json"
	"testing"
)

func TestMap(t *testing.T) {
	newMap := map[string]string{"abc": "abc"}

	t.Log(newMap)
	jsonString, err := json.Marshal(newMap)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(jsonString))

	// make 只能用于创建  map 、 slice 、 是声明变量且初始化
	mapCreated := make(map[string]float32)

	mapCreated["abc"] = 3.1415
	mapCreated["bcd"] = 5.26

	for k, v := range mapCreated {
		t.Log(k, "+++", v)
	}

	var strings []string
	for k := range mapCreated {
		t.Log("key 为 ", k)
		strings = append(strings, k)
	}
	t.Log(strings)

}
