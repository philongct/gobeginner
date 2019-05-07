package emptyinterface

import (
	"fmt"
)

func processMap(parentPath string, m map[string]interface{}) []map[string]interface{} {
	var queue []map[string]interface{}
	for k, v := range m {
		queue = append(queue, map[string]interface{}{parentPath + k: v})
	}

	return queue
}

func processSlice(parentPath string, arr []interface{}) []map[string]interface{} {
	var queue []map[string]interface{}
	path := parentPath[:len(parentPath)-1]
	for i, e := range arr {
		queue = append(queue, map[string]interface{}{fmt.Sprintf("%s[%d]", path, i): e})
	}

	return queue
}

func Flatten(m map[string]interface{}, sep string) map[string]interface{} {
	ret := make(map[string]interface{})

	queue := processMap("", m)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// the map taken from queue node has only 1 key and 1 value:
		// the key is the path to the value
		for k, v := range node {
			switch vt := v.(type) {
			case map[string]interface{}:
				queue = append(queue, processMap(k+sep, vt)...)
			case []interface{}:
				queue = append(queue, processSlice(k+sep, vt)...)
			default:
				ret[k] = v
			}
		}
	}

	return ret
}
