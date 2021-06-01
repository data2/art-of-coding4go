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

	var start = 0
	for start < 10 {
		start++
		fmt.Println(start)
	}

	if !isOk {
		fmt.Println("不ok")
	}

}

func sum(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func arr() {
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr1[2])

	arr1 = [3]int{0: 2, 2: 2}
	fmt.Println(arr1)

}

func point() {
	var a = 10
	var point *int = &a

	fmt.Println(fmt.Sprintf("a变量的指针地址为%x", point))
	fmt.Println(fmt.Sprintf("a变量的指针地址为%x", &a))

	fmt.Println(fmt.Sprintf("a变量值%d", *point))

}
