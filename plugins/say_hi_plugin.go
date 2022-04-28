package main

import "fmt"

var Hello string //必须是大写，小写的时候提示找不着此 symbol

//加载插件时会被调用
func init()  {
	fmt.Println("Test plugin init")
}

func SayFunc()  {
	fmt.Printf("Say func will print %s %s\n", Hello, "World！") //必须是大写，小写的时候提示找不着此 symbol
}
