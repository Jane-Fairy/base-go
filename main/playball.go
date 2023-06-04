package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func playerDemo(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("end , palyer : %s won!", name)
			return
		}
		n := rand.Intn(10000)
		if n%14 == 0 {
			fmt.Printf("palyer : %s miss!", name)
			close(court)
			return
		}
		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手
		court <- ball
	}
}

func playball() {
	court := make(chan int)
	// 计数加 2，表示要等待两个goroutine
	wg.Add(2)
	// 启动两个选手
	go playerDemo("Nadal", court)
	go playerDemo("Djokovic", court)
	// 发球
	court <- 1
	// 等待游戏结束
	wg.Wait()
}
