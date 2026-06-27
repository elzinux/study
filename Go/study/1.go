package main

import "fmt"

func print1() {
	fmt.Println("Hello base.go!")
	fmt.Printf("%v, %v, %v\n", 0, 0.0, false)

	x := 1
	a := make([]int, 10, 20)
	b := map[string]int{
		"123": 123,
		"234": 456,
	}
	fmt.Println(x, a, len(a), cap(a)) //1 [0 0 0 0 0 0 0 0 0 0] 10 20
	fmt.Println(b)                    //map[123:123 234:456]

}

func arr_map() {
	// defind array
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr1[0] = 10
	fmt.Println(arr1)

	//defind map
	// var mp1 map[string]int
	mp1 := make(map[string]int)
	mp1["123"] = 123

	mp2 := map[string]int{}
	mp2["123"] = 123

	mp3 := map[string]int{
		"123": 123,
	}
	mp3["234"] = 234

	// get value
	v1, f1 := mp1["123"]
	v2, f2 := mp1["234"]
	if f1 {
		fmt.Println(v1) // 123
	}
	if f2 {
		fmt.Println(v2) // nothing
	}

	// defind set
	set := map[int]struct{}{}
	set[1] = struct{}{}
	set[2] = struct{}{}
	delete(set, 2)
	for set_v := range set {
		fmt.Print(set_v)
	} // 1
}

// 用于屏蔽格式化，可以让格式化的时候不检查某一个函数或者某一行代码
// 只能屏蔽自动检查
//
//nolint:gosimple
func for_() {
	arr := make([]int, 10)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()

	i := 0
	for i < len(arr) {
		fmt.Print(arr[i])
		i++
	}
	fmt.Println()

	for i := range len(arr) {
		fmt.Print(arr[i])
	}
	fmt.Println()
	for i := range arr {
		fmt.Print(arr[i])
	}
	fmt.Println()
	for i, v := range arr {
		fmt.Print(arr[i], v, "|")
	}
	fmt.Println()
	mp := map[string]int{
		"123": 123,
	}
	for key, val := range mp {
		fmt.Println(key, val)
	}
}

func diff(a, b int) bool {
	if a != b {
		return true
	}
	return false
}

// 泛型函数的两种定义方法
func diff1[T int | float32 | string](a, b T) bool {
	if a != b {
		return true
	}
	return false
}

type T1 interface {
	int | float32 | string
}

func diff2[T T1](a, b T) bool {
	if a != b {
		return true
	}
	return false
}

func main1() {
	// print1()
	// arr_map()
	// for_()

	x := max(10, 11)
	fmt.Println(x, diff(10, 11)) // 11 true

}
