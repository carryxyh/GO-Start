package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string //这里如果世ServerName 解析正常，若是serverName，失败，但是我们看到json中的键都是小写字母开头
	//这里有个技巧，当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	//以下是解析不知道结构的json，我们这里用interface
	var f interface{}
	err := json.Unmarshal([]byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`), &f)
	if err == nil {
		fmt.Println(err)
	}
	//map[Name:Wednesday Age:6 Parents:[Gomez Morticia]] 输出
	fmt.Println(f)

	//访问数据,通过断言
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
