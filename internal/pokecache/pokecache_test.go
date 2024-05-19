package pokecache

import (
	"testing"
	"time"
)


func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

// func TestAddGetCache(t *testing.T) {
// 	cache := NewCache()

// 	cache.Add("key1", []byte("val1"))
// 	actual, ok := cache.Get("key1")
// 	if !ok {
// 		t.Error("key1 not found")
// 	}
// 	if string(actual) != "val1" {
// 		t.Error("value doesn't match")
// 	}
// }

/*====  重构测试 & 表驱动测试 ======*/
func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct{
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for _, c := range cases {
		cache.Add(c.key, c.val)
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("%s is not found", c.key)
			return
		}
		if string(actual) != string(c.val) {
			t.Errorf("%s doesn't match %s", actual, c.val)
			return
		}
	}

}


func TestReap(t *testing.T) {
	interval := time.Millisecond * 10 // 10ms
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond) // 11ms
	
	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}

}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10 // 10ms
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2) // 11ms
	
	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}

}