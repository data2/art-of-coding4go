package main

import "fmt"

func basic() {

	var url = "你好，%s"
	var realUrl = fmt.Sprintf(url, "go")
	fmt.Println(realUrl)

	var number1 int
	var isOk bool
	fmt.Println(fmt.Sprintf("数值默认值为%d, 布尔默认值为%t", number1, isOk))

	number1 = 2
	var number2 = 3
	fmt.Println(fmt.Sprintf("2 * 3 = %d", number1*number2))

	var float1 float64 = 2.1
	fmt.Println(float1)

	str := "你好"
	fmt.Println(str)

	var i, j = 10, 20
	fmt.Println(i * j)

}
