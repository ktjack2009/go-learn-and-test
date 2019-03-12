/*
	1. 布尔值转化为json后还是布尔类型
	2. 浮点数和整型转化为常规数字
	3. 字符串将以utf-8输出为unicode
	4. 数组和切片会转化为json里面的数组
	5. []byte类型的值会被转化为Base64编码后的字符串
	6. 结构体会转化为json对象，并且只有结构体里面以大写字母开头的可被导出的字段才会被转化输出
	7. 转化一个map类型的数据结构时，该数据的类型必须是map[string]T
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonData struct {
	Is    bool              `json:"is"`
	Name  string            `json:"name"`
	Num   float64           `json:"num"`
	Slice []string          `json:"slice"`
	Byte  []byte            `json:"byte"`
	Map   map[string]string `json:"map"`
	I     interface{}
}

func main() {
	//JsonDemo1()
	//JsonDemo2()
	//JsonDemo3()

	// 流式读写
}

func JsonDemo1() {
	obj := JsonData{
		Is:    true,
		Name:  "dsj",
		Num:   10.9,
		Slice: []string{"dsj", "wrc"},
		Byte:  []byte("hello"),
		Map:   map[string]string{"1": "1", "2": "2"},
		I:     string([]byte("hello")),
	}
	data, _ := json.Marshal(obj)

	file, _ := os.Create("no.json")
	file.Write(data)
	file.Close()
}

func JsonDemo2() {
	// 推荐使用ioutil读取文件
	data, _ := ioutil.ReadFile("no.json")
	var obj JsonData
	json.Unmarshal(data, &obj)
	fmt.Println(obj)
}

func JsonDemo3() {
	data, _ := ioutil.ReadFile("no.json")
	obj := make(map[string]interface{})
	json.Unmarshal(data, &obj)
	fmt.Println(obj)
}
