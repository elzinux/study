package main

import (
	"bufio"
	"fmt"
	"os"
)

func read() {
	dir := "./test.txt"
	// 一次性读取
	byteData, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println("fail")
	}
	fmt.Println(byteData) // 打印一个整数列表，类似：[72 101 108 108 111 32 87 111 114 108 100 33]
	fmt.Println(string(byteData)) // 打印文件中的内容，类似：Hello World!

	// 按行读取
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("打开文件的时候发生了错误！！！")
	}
	defer file.Close()

	// buf := bufio.NewReader(file)
	// for {
	// 	line, _, err := buf.ReadLine()
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(string(line))
	// 	fmt.Println("---")
	// }

	scan := bufio.NewScanner(file)
	// scan.Split(bufio.ScanBytes) // 按字节读，中文等会乱码
	scan.Split(bufio.ScanRunes) // 几乎能读所有字符
	// scan.Split(bufio.ScanWords) // 按照单词读，只分割空白符号
	// scan.Split(bufio.ScanLines) // 按照行读 

	for scan.Scan() {
		fmt.Println(scan.Text())
		fmt.Println("-")
	}
	if err := scan.Err(); err != nil {
		fmt.Println("读取文件的时候发生了错误！！！")
	}
}

func write() {
	dir := "./test.txt"
	str := "我是你！！！"
	// os.ModePerm 打印出来是 8 进制 0777 所有者，同组用户，其他用户，都拥有“读，写，执行”的权利
	err1 := os.WriteFile(dir, []byte(str), os.ModePerm)
	if err1 != nil {
		fmt.Println("写内容的时候出错了，请检查！！！", err1)
	}

	file, err := os.OpenFile(dir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 4)
	if err != nil {
		fmt.Println("打开文件出错了！！！")
	}
	defer file.Close()

	_, err2 := file.WriteString(str + "追加的")
	if err2 != nil {
		fmt.Println("追击内容出现了问题！！！", err2)
	}

}

func main() {
	// read()
	// write()
	fmt.Println(os.ModePerm)
}
