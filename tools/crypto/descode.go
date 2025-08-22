package crypto

import (
	"encoding/json"
	"strings"
)

func JsonDescodeToMap(jsonstr string) map[string]interface{} {
	// 使用NewDecoder函数创建一个解码器
	decoder := json.NewDecoder(strings.NewReader(jsonstr))
	// 逐步解码JSON数据
	var data map[string]interface{}
	for decoder.More() {
		if err := decoder.Decode(&data); err != nil {
			panic(err)
		}
	}
	return data
}

func JsonDecode(jsonStr string, data interface{}) error {
	return json.Unmarshal([]byte(jsonStr), data)
}
