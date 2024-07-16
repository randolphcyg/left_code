package main

import (
	"fmt"
	"math/rand"
)

/*
设计RandomPool结构
*/

type Pool struct {
	KeyIndexMap map[string]int
	IndexKeyMap map[int]string
	Size        int
}

func InitPool() *Pool {
	return &Pool{
		KeyIndexMap: make(map[string]int),
		IndexKeyMap: make(map[int]string),
		Size:        0,
	}
}

func (p *Pool) insert(key string) {
	if _, ok := p.KeyIndexMap[key]; !ok {
		p.KeyIndexMap[key] = p.Size
		p.IndexKeyMap[p.Size] = key
		p.Size++
	}
}

func (p *Pool) delete(key string) {
	if _, ok := p.KeyIndexMap[key]; ok {
		deleteIndex, _ := p.KeyIndexMap[key]
		p.Size--
		lastIndex := p.Size
		lastKey, _ := p.IndexKeyMap[lastIndex]
		p.KeyIndexMap[lastKey] = deleteIndex
		p.IndexKeyMap[deleteIndex] = lastKey
		delete(p.KeyIndexMap, key)
		delete(p.IndexKeyMap, lastIndex)
	}
}

func (p *Pool) getRandom() string {
	if p.Size == 0 {
		return ""
	}
	randomIndex := rand.Intn(p.Size)
	return p.IndexKeyMap[randomIndex]
}

func main() {
	pool := InitPool()
	pool.insert("ke")
	pool.insert("nan")
	pool.insert("love")
	pool.insert("ran")
	pool.insert("yuan")
	fmt.Println(pool.getRandom())
	pool.delete("ran")

	fmt.Println(rand.Intn(4))

	for i := 0; i < 5; i++ {
		fmt.Println(i, pool.getRandom())
	}
}
