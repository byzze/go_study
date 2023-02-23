package singleflight

import (
	"log"
	"sync"
)

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call
}

// 应对并发请求时
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		log.Println("Do....")
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait() // 后来的协程等待获取源数据放入缓存
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn() // 先来的协程获取数据放入缓存
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key) // 删除数据, 防止缓存数据是旧值
	g.mu.Unlock()
	log.Println("End....")
	return c.val, c.err
}