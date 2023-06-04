package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		//v[i] += u.Op()
	}
}

func parallelComputing() {

}

func syncGroup() {
	var group sync.WaitGroup
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	for _, url := range urls {
		group.Add(1)
		go func(url string) {
			// 使用defer, 表示函数完成时将等待组值减1

			// 使用http访问提供的地址
			_, err := http.Get(url)
			// 访问完成后, 打印地址和可能发生的错误
			fmt.Println(url, err)
			// 通过参数传递url地址
			group.Done()
		}(url)
	}

	group.Wait()
	fmt.Println("over")
}

func s() {

}
