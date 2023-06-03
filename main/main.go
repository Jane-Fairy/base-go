package main

import (
	"bufio"
	"bytes"
	"fmt"
	"learn_golang/middleware"
	_ "learn_golang/middleware"
	"log"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"time"
)

type interfaceDemo interface {
	Call(interface{})
}

type Strut struct {
}

func (s *Strut) Call(p interface{}) {
	fmt.Println(p)
}

func visit(list []int, f func(int)) {
	for i := range list {
		f(i)
	}
}

func fibonacci(n int) (res int) {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

type funcDemo func(interface{})

func (f *funcDemo) Call(p interface{}) {
	fmt.Println(p)
}

func Factorial(n int) (res int) {
	if n <= 1 {
		return 1
	}
	return fibonacci(n) * fibonacci(n-1)
}

func player(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

func GetAare(weight int, height int) int {
	return weight * height
}

func ProtectRun(entry func()) {
	now := time.Now()

	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("errorr")
		}
	}()

	time.Since(now)
}

type inners struct {
	in1  int
	inl2 int
}

type outers struct {
	inner inners
	b     int
	c     float32
	int
	inners
}

type Road int

func finalizerDeo(r *Road) {
	log.Println("road", *r)
}

type myStringSort []string

func (m myStringSort) Len() int {
	return len(m)
}

func (m myStringSort) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m myStringSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func init() {
	fmt.Println("this is main.")
}

func main() {

	middleware.Middle{}.LoggerToFile()

	names := myStringSort{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用sort包进行排序
	sort.Sort(names)
	// 遍历打印结果

	//i := int8(9)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
	//Read 方法
	date := []byte("hello world")
	reader := bytes.NewReader(date)

	newReader := bufio.NewReader(reader)
	var buf [128]byte
	newReader.Read(buf[:])

	var rd Road = Road(999)
	r := &rd
	runtime.SetFinalizer(r, finalizerDeo)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}

	go func() {
		//fibonacci
		fmt.Println("斐波那契：")
		fmt.Println(fibonacci(10))

		//阶层
		fmt.Println("n的阶层")
		fmt.Println(Factorial(2))

		var invoker interfaceDemo
		s := new(Strut)

		invoker = s
		invoker.Call("hello struct")
	}()

	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}

	for true {

	}
	//fmt.Println("ss")
	//for y := 1; y <= 9; y++ {
	//	for x := 1; x <= y; x++ {
	//		fmt.Printf("%d x %d = %d \t", x, y, x*y)
	//
	//	}
	//	goto mark
	//	fmt.Println()
	//}

	//mark:
	//	fmt.Println("hello")
	//二维切片
	//var slice [][]int
	//slice = [][]int{{10}, {100, 200}}
	//
	////链表
	//list1 := list.New()
	//var list2 list.List
	//list1.PushBack(1)
	//list1.PushFront(1)
	//list1.InsertAfter(1, list1.PushBack(1))
	//
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

func read(n int, cn chan struct{}) {
	rw.RLock()
	fmt.Println("goroutine %d 开始读取\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束 ， 值为： %d\n ", n, v)
	rw.Unlock()
	cn <- struct{}{}
}

func write(n int, cn chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	cn <- struct{}{}
}

var count int

var rw sync.RWMutex