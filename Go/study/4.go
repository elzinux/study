package main

import (
	"errors"
	"fmt"
	"sync"
)

// 使用 lock 来锁住资源，互斥锁
func add(x *int, wait *sync.WaitGroup, lock *sync.Mutex) {
	defer wait.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		*x++
		lock.Unlock()
	}
}

func reduce(x *int, wait *sync.WaitGroup, lock *sync.Mutex) {
	defer wait.Done()
	for i := 0; i < 10000; i ++ {
		lock.Lock()
		*x--
		lock.Unlock()
	}
}


// 错误处理
var ERROR0 = errors.New("被除数不能为 0 !!!")

func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, errors.New("除数不能为 0 !!!")
		return 0, fmt.Errorf("%v是除数，他不能为 0 !!!", b)
	} else if a == 0 {
		return 0, ERROR0
	} else {
		return a / b, nil
	}
}

func test_panic() {
	fmt.Println("befor panic")
	// defer func () {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("test_panic 发生了 panic !!!")
	// 	}
	// }()

	x, y := 10, 0; fmt.Println(x / y)
	// 发生 panic 之后的代码都不回执行
	fmt.Println("after panic")
}

func test_panic1() {
	// 捕获了子函数中的 panic
	defer func () {
		if err := recover(); err != nil {
			fmt.Println("test_panic 发生了 panic !!!")
		}
	}()
	test_panic()
	// 发生 panic 之后，父函数中的后续代码也不回执行
	fmt.Println("after test_panic1")

}

func main4() {
	// var x int
	// wait := sync.WaitGroup{}
	// // 同步锁，抢到的协程执行，其他全部阻塞
	// // 阻塞的协程不占用 cpu
	// lock := sync.Mutex{}

	// wait.Add(2)
	// go add(&x, &wait, &lock)
	// go reduce(&x, &wait, &lock)
	// wait.Wait()
	// fmt.Println(x)

	// 错误处理
	// res, err := divide(0, 3)
	// if err == nil {
	// 	fmt.Println(res)
	// } else if errors.Is(err, ERROR0){
	// 	fmt.Println("我们的除法要求被除数也不能为 0 !!!")
	// } else {
	// 	fmt.Println(err) // 0 是除数，他不能为 0 !!!
	// }

	test_panic1()
	// 发生 panic 之后 mian 函数中的后续代码还可以正常执行
	fmt.Println("main after panic !!!")
}
