package util

func AppendMap(map1, map2 map[string]interface{}) map[string]interface{} {
	mp3 := make(map[string]interface{})
	for k,v := range map1 {
		if _,ok := map1[k]; ok {
			mp3[k] = v
		}
	}

	for k,v := range map2 {
		if _,ok := map2[k]; ok {
			mp3[k] = v
		}
	}

	return mp3
}
