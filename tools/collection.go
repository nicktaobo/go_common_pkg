package tools

import "reflect"

func MergeMap[V any](src map[string]V, dst map[string]V) map[string]V {
	for range src {
		for k2, v2 := range dst {
			// 不存在则合并
			if _, b := src[k2]; !b {
				src[k2] = v2
			}
		}
	}
	return src
}

// StructToMap will change struct to map
func StructToMap(inter interface{}) map[string]interface{} {
	param := make(map[string]interface{})

	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	for i := 0; i < t.NumField(); i++ {
		param[t.Field(i).Name] = v.Field(i).Interface()
	}
	return param
}
