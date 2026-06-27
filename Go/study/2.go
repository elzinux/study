package main

import (
	// "GoStudy/test"
	"fmt"
)

// 自定义别名
type i64 int64

// 一个文件夹下的所有 go 文件只能有一个 main 函数
// func test_import() {
// 	fmt.Println("Hello mid.go!")
// 	test.Print() // 调用 test 中的函数，需要首字母大写才能被导入使用
// 	print1()     // 直接使用 base.go 中的函数，需要 go run high.go base.go，否则报错
// }

// init 会在包导入的时候按顺序执行（类似队列），不能传参数，不能返回值
// 先打印 init 1 后打印 init 2
// 运行直接执行，不需要调用
func init() { fmt.Print("init 1 ") }
func init() { fmt.Println("init 2") }

func defer1() {
	// defer 的执行是在函数结束后，返回前执行，并且是逆序执行（类似栈）
	// 先打印 defer 2 后打印 defer 1
	defer fmt.Println("defer 1")
	defer func() {
		fmt.Println("defer 2")
	}()
}

// ----- 面向对象 -----
// 结构体
type People struct {
	Name string
	Age  int
}

// 为结构体绑定方法
func (p People) printName() {
	fmt.Println("name:", p.Name)
}

type Student struct {
	People   // 继承
	stuId    int
	stuClass int
}

// 这是实现了 Student 的值方法
// 调用的时候直接拷贝完整的 Student
func (s Student) printId() {
	fmt.Println("stuID:", s.stuId)
}

// 这是实现 Student 的指针方法
// 调用的时候只有 8 字节的内存占用
// 结构体大的时候尽量使用指针绑定方法
func (s *Student) printClass() {
	fmt.Println("Class:", s.stuClass)
}

func setId(s Student, id int) {
	s.stuId = id
}

func setClass(s *Student, class int) {
	s.stuClass = class
}

// ----- 接口 -----
// 接口是一组仅包含 方法名，参数，返回值 的未具体实现的方法集合
// 接口本身不能绑定方法，接口是值类型，保存的是：值 + 原始类型
type Animal interface {
	eat()
	run()
}

// 结构体实现了接口中的所有方法，结构体就实现了该接口
type Cat struct {
	Name string
	cat  string
}

func (c Cat) eat() {
	fmt.Println(c.Name, "在吃东西")
}

func (c Cat) run() {
	fmt.Println(c.Name, "在跑路")
}

// Dog 实现接口
type Dog struct {
	Name string
	dog  string
}

func (d *Dog) eat() {
	fmt.Println(d.Name, "在吃东西")
}

func (d *Dog) run() {
	fmt.Println(d.Name, "在跑路")
}

// 通过接口可以直接调用方法
func eat(obj Animal) {
	obj.eat()
}

func run(obj Animal) {
	// 只有在 seitch 中可以使用 .(type)
	// 并且只有 interface 可以使用
	switch obj.(type) {
	case Cat:
		fmt.Println("我是一只猫，在跑路")
	case *Dog:
		fmt.Println("我是一只狗，在跑路")
	}
}

func main2() {
	// var x i64 // 即使和 int64 的变量比较也需要转换
	// fmt.Printf("%T", x) //main.i64

	// test_import()
	// defer1()

	// s := Student{
	// 	stuId: 123,
	// 	People: People{
	// 		Name: "123",
	// 		Age:  123,
	// 	},
	// }

	// s.printId()
	// s.printName() // s.People.printName()，当 Student 没有实现printName时两种写法等价
	// s.printClass()

	// setId(s, 321) // 修改不成功
	// fmt.Println(s.stuId) // 打印 123
	// setClass(&s, 321) // 修改成功
	// fmt.Println(s.stuClass) // 打印 321

 	// 接口
	var c Animal
	c = Cat{ Name: "maomao", }
	// 上下两种初始化都可以
	// c = &Cat{ Name: "maomao", }
	// c.eat() // maomao 在吃东西
	// c.run() // maomao 在跑路

	d := Dog{
		Name: "gougou",
	}

	var dog Animal
	dog = &Dog{ Name: "123", }
	// 只能使用上面的方法初始化，因为dog的值类型没有实现接口 Animal
	// dog = Dog{ Name: "123", }
	dog.eat()

	// eat(c) // maomao 在吃东西
	// eat(d) // gougou 在吃东西
	run(c) // 我是一只猫，在跑路
	run(&d) // 我是一只狗，在跑路
}
