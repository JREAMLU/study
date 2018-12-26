package jsync

import (
	"sync"
)

type Cache struct {
	lock sync.RWMutex
	data []int // 实际数据比这个复杂很多有很多维度
}

func (c *Cache) Get() []int {
	c.lock.RLock()
	defer c.lock.RUnlock()

	var res []int

	// 筛选数据， 简单写一个筛选过程
	for i := range c.data {
		if c.data[i] > 10 {
			res = append(res, c.data[i])
		}
	}

	return res
}

func (c *Cache) GetC(next chan struct{}) chan int {
	ch := make(chan int, 1)

	go func() {
		c.lock.RLock()
		defer c.lock.RUnlock()
		defer close(ch)

		// 筛选数据， 简单写一个筛选过程
		for i := range c.data {
			if c.data[i] > 10 {

				ch <- c.data[i]

				if _, ok := <-next; !ok {
					return
				}
			}
		}
	}()

	return ch
}
