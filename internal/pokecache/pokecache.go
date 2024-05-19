package pokecache

import (
	"sync"
	"time"
)

// 缓存条目
type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{}, // 如果使用了 *sync.Mutex, 需要手动初始化
		//可能导致 nil pointer dereference 错误，如果指针没有正确初始化。
	}
	// 这里如果不加 go, 会被阻塞
	/* 由于 reapLoop 方法中的循环会无限执行，
	因此必须在一个新的 goroutine 中运行，
	否则会阻塞主 goroutine，导致程序无法继续执行其他任务：
	*/
	/* 加 go 以后启动了一个新的 goroutine，
	确保 reapLoop 在后台运行，而不会阻塞主 goroutine。
	*/
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 此时是 cacheE 是 map 的 value
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}


// 移除过期的缓存条目

func (c *Cache) reapLoop(interval time.Duration) {
	// time.Ticker 是一个计时器，
	// 它会每隔指定的时间间隔发送一个事件到它的通道 C。
	ticker := time.NewTicker(interval)
	// for range ticker.C 的意思是：每当通道上有新的事件时，执行一次循环体
	for range ticker.C { // ticker.C 是一个通道，每当时间间隔过去时，通道上就会发送一个事件。
		c.reap(interval)
	}
}

/*
 1) time.Now() 获取本地时间
 2) .UTC() 将当前时间转换为协协调世界时 (UTC)。
 3) .Add(-interval) 将时间向前移动指定的间隔。
由于 interval 是正值，加上负号表示将时间向过去移动。
*/
func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 获取 interval 时间之前的 UTC 时间
	timeAgo := time.Now().UTC().Add(-interval)

	for k, v := range c.cache {
		// 如果创建时间早于 timeAgo
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}