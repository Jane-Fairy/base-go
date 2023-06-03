package main

import "sync"

/*
sync.Pool 是临时对象池，存储的是临时对象，不可以用它来存储 socket 长连接和数据库连接池等。

sync.Pool 本质是用来保存和复用临时对象，以减少内存分配，降低 GC 压力，比如需要使用一个对象，就去 Pool 里面拿，如果拿不到就分配一份，这比起不停生成新的对象，用完了再等待 GC 回收要高效的多。
*/

type student struct {
	age int
}

var studentPool = &sync.Pool{
	New: func() interface{} {
		return new(student)
	},
}

type Student struct {
	Name interface{}
}

func New(name string, age int) *student {
	return studentPool.Get().(*student)
}
