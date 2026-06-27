package main

import (
	"fmt"
	"sync"
	"time"
)

func say() {
	fmt.Println("开始说话")
	time.Sleep(1 * time.Second)
	fmt.Println("说话结束")
}

var wait = sync.WaitGroup{}

// 一个文件夹下的所有文件之中不能有同名字的函数
// 否则 go run high.go mid.go 会报错
// 但是直接运行一个 go run high.go 不会出错
func eating() {
	fmt.Println("开始吃东西")
	fmt.Println("吃东西结束")
	wait.Done()
}

var moneyChan = make(chan int)
var nameChan = make(chan string)

func pay(name string, money int, wait *sync.WaitGroup) {
	defer wait.Done()
	fmt.Println(name, "开始买东西")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "结束买东西")
	moneyChan <- money 
	nameChan <- name
}

func main3() {
	fmt.Println("Hello high.go!")
	// 协程，执行顺序是完全随机的
	// go say()
	// go say()
	// // 此处如果不延时 say 不会之行，因为 main 之行完了，资源被全部回收了，协程无法完成执行
	// time.Sleep(2 * time.Second) 

	// wait.Add(2)
	// go eating()
	// go eating()
	// wait.Wait()
	// fmt.Println("high.go 主程序结束")
	// 打印内容如下：
	// 开始吃东西
	// 吃东西结束
	// 开始吃东西
	// 吃东西结束
	// high.go 主程序结束

	// var c chan int 
	// 初始化一个带缓冲的通道，如果 make 中不给缓冲则为 0
	// c := make(chan int, 1)
	// c <- 1

	// c <- 2 // 报错，因为缓冲区装满了，必须有代码接走信息才能运行，否则 deadlock
	// <- c // 即使取值也会报错

	// x, ok := <- c // 只有不超过满缓冲区容量才不会报错
	// fmt.Println(x, ok) // 1 true

	// fmt.Println(c) // 打印的是一个地址
	// close(c)

	// 协程与chan
	var wait sync.WaitGroup
	time1 := time.Now()
	wait.Add(2)
	go pay("maomao", 10, &wait)
	go pay("gougou", 20, &wait)

	go func() {
		defer close(moneyChan)
		defer close(nameChan)
		wait.Wait()
		fmt.Println("我是辅助协程，等待业务协程执行完毕，我将关闭所有通道！！！")
	}()

	// 用于从单个 chan 中取值
	// moneyList := []int{}
	// for money := range moneyChan {
	// 	moneyList = append(moneyList, money)
	// }

	moneyList := []int{}
	nameList := []string{}

	endTimeChan := time.After(2 * time.Second)

	event := func() {
		for {
			select {
			case money, ok := <-moneyChan:
				if !ok {
					return 
				}
				moneyList = append(moneyList, money)
			case name, ok := <-nameChan:
				if !ok {
					return 
				}
				nameList = append(nameList, name)
			case <-endTimeChan:
				fmt.Println("超时关店退出了！！！")
				return
			}
		}
	}

	event()

	fmt.Println("超市关门了", time.Since(time1)) // 打印 1s 多一点，即使每个 pay 都需要 1s 执行
	fmt.Println("收到的钱：", moneyList)
	fmt.Println("客户名字：", nameList)
	// 代码执行逻辑，event 中的死循环防止 main 提前执行完后断掉所有协程

}