package main

import (
	"fmt"
	"strconv"
	"sync"
)

const count = 1000

func main() {
	cacheMap := NewCacheMap()

	var wg sync.WaitGroup
	wg.Add(count * 2)

	// проверка на паралельную запись
	for i := 0; i < count; i++ {
		val := strconv.Itoa(i)

		go func() {
			defer wg.Done()
			cacheMap.Set(val, val)
		}()

	}

	// паралельное чтение
	for i := 0; i < count; i++ {
		val := strconv.Itoa(i)

		go func() {
			defer wg.Done()
			fmt.Println(cacheMap.Get(val))
		}()
	}

	// по сути "тесты" сверху одновременно и записывают
	wg.Wait()

}

type cacheMap struct {
	mp map[string]string
	m  sync.Mutex
}

func NewCacheMap() *cacheMap {
	return &cacheMap{
		mp: make(map[string]string),
	}
}

func (m *cacheMap) Set(key, value string) {
	m.m.Lock()
	defer m.m.Unlock()

	m.mp[key] = value
}

func (m *cacheMap) Get(key string) (string, error) {
	m.m.Lock()
	defer m.m.Unlock()

	if val, ok := m.mp[key]; !ok {
		return "", fmt.Errorf("key does not exists")
	} else {
		return val, nil
	}
}
