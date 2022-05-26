package main

import (
	"fmt"
	"reflect"
)

type GiftPacket struct {
	PushType       string   `json:"pushtype"`
}

func main() {
	fmt.Println("start")
	h := GiftPacket{}
	refH := reflect.TypeOf(h)
	//fmt.Println(refH.PkgPath())
	//fmt.Println(refH.Kind())
	//fmt.Println(refH.Name())
	//fmt.Println(refH.String())
	////创建一个结构体
	//refH = refH.Elem()
	////获取结构体名
	//fmt.Println(refH.Name())
	////获取类别
	//fmt.Println(refH.Kind())
	for i := 0; i < refH.NumField(); i++ {
		field := refH.Field(i)
		//fmt.Println("结构体里的字段名",field.Name)
		//fmt.Println("结构体里的字段属性:",field.Type)
		//fmt.Println("结构体里面的字段的tag标签",field.Tag)
		if jsonName, ok := field.Tag.Lookup("json"); ok {
			fmt.Print("\n\"", jsonName, "\": ")
		} else {
			fmt.Println("不是json")
			return
		}
		//fmt.Print(field.Type.String())
		switch field.Type.String() {
		case "string":
			fmt.Print("\"\"")
		case "int":
			fmt.Print("\"\"")
		case "[]string":
			fmt.Print("[\"\", \"\"]")
		case "[]int":
			fmt.Print("[0, 0, 0]")
		default:
			fmt.Print("\"未知类型\"")
		}
		fmt.Print(",")
	}
	//提取tag标签里的信息
	name, bool := refH.FieldByName("Name")
	if bool {
		fmt.Println("tag标签的信息为",name.Tag.Get("json"))
	}

}
