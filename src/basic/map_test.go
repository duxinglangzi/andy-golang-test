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

// 合并
func TestMergeMap(t *testing.T) {
	newMap := make(map[string]interface{})
	newMap["a"] = "sssss"
	newMap["b"] = 123
	t.Log(newMap)

	oldMap := make(map[string]interface{})
	oldMap["a"] = "sssss"
	oldMap["b"] = 123
	oldMap["c"] = "ahhahahahahhahah"

	//m := mergeMap(newMap, oldMap)


}

func mergeMap(a map[interface{}]interface{}, b map[interface{}]interface{}) map[interface{}]interface{} {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	for k, v := range a {
		_, containsKey := b[k]
		if !containsKey {
			b[k] = v
		}
	}
	return b
}
