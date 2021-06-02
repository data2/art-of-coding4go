package basic

import "fmt"

type Person struct {
	name   string
	age    int
	height int
}

type Dog struct {
}

type Walk interface {
	walk()
}

func Test() {

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

func Sum(a, b int) int {
	return a + b
}

func Swap(a, b string) (string, string) {
	return b, a
}

func Arr() {
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr1[2])

	arr1 = [3]int{0: 2, 2: 2}
	fmt.Println(arr1)

	slice := make([]int, 10, 12)
	slice[0] = 1
	fmt.Println(fmt.Sprintf("切片大小=%d, 切片容量=%d", len(slice), cap(slice)))
	fmt.Println(slice)

	slice = append(slice, 11, 12)
	fmt.Println(fmt.Sprintf("切片大小=%d, 切片容量=%d", len(slice), cap(slice)))
	fmt.Println(slice)

	slice = append(slice, 13)
	fmt.Println(fmt.Sprintf("切片大小=%d, 切片容量=%d", len(slice), cap(slice)))
	fmt.Println(slice)

	slice2 := make([]int, 11, 11)
	copy(slice2, slice)
	fmt.Println(slice2)

}

func Map() {
	map1 := make(map[string]string)
	map1["china"] = "beijing"
	map1["usa"] = "华盛顿"
	for country := range map1 {
		fmt.Println(fmt.Sprintf("国家=%s，首都=%s", country, map1[country]))
	}

	value, ok := map1["china"]
	if ok {
		fmt.Println("china首都" + value)
	} else {
		fmt.Println("map不存在china")
	}

	value1, ok1 := map1["japan"]
	if ok1 {
		fmt.Println("japan首都" + value1)
	} else {
		fmt.Println("map不存在japan")
	}

	delete(map1, "china")
	fmt.Println(map1)

}

func Point() {
	var a = 10
	var point *int = &a

	fmt.Println(fmt.Sprintf("a变量的指针地址为%x", point))
	fmt.Println(fmt.Sprintf("a变量的指针地址为%x", &a))

	fmt.Println(fmt.Sprintf("a变量值%d", *point))

}

func TestStruct() {
	p1 := Person{name: "zhangsan", height: 178}
	fmt.Println(p1)

	var p2 Person = Person{"李四", 20, 180}
	fmt.Println(p2.name)

	p2.age += 1
	p2Point := &p2
	fmt.Println(p2Point.age)

}

func (person Person) walk() {
	fmt.Println("人 - 走路")
}

func (dog Dog) walk() {
	fmt.Println("狗 - 跑")
}

func TestInterface() {
	var p Person = Person{name: "某某人"}
	var dog Dog = Dog{}
	p.walk()
	dog.walk()
}
