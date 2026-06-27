package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)
var nameChan = make(chan string)

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Println(name, "开始买东西")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "结束买东西")
	moneyChan <- money 
	nameChan <- name
	wait.Done()
}

func main() {
	fmt.Println("Hello World!")
		var wait sync.WaitGroup
	time1 := time.Now()
	wait.Add(2)
	go pay("maomao", 10, &wait)
	go pay("gougou", 20, &wait)

	go func() {
		fmt.Println("关闭协程！！！")
		defer close(moneyChan)
		defer close(nameChan)
		wait.Wait()
	}()

	// 用于从单个 chan 中取值
	// moneyList := []int{}
	// for money := range moneyChan {
	// 	moneyList = append(moneyList, money)
	// }

	moneyList := []int{}
	nameList := []string{}

	event := func() {
		for {
			select {
			case money, ok := <-moneyChan:
				if ok {
					moneyList = append(moneyList, money)
				}
			case name, ok := <-nameChan:
				if ok {
					nameList = append(nameList, name)
				}
			}
		}
	}

	event()

	fmt.Println("超市关门了", time.Since(time1)) // 打印 1s 多一点，即使每个 pay 都需要 1s 执行
	fmt.Println("收到的钱：", moneyList)
	fmt.Println("客户名字：", nameList)
}
