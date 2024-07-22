package main

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

/*
一致性哈希
*/

// Node 表示一个缓存节点
type Node struct {
	IP   string
	Hash uint32
}

// ConsistentHashRing 表示一致性哈希环
type ConsistentHashRing struct {
	Nodes []Node
}

// NewConsistentHashRing 创建一个新的一致性哈希环
func NewConsistentHashRing() *ConsistentHashRing {
	return &ConsistentHashRing{}
}

// AddNode 添加一个新的缓存节点
func (chr *ConsistentHashRing) AddNode(ip string) {
	hash := hash(ip)
	node := Node{IP: ip, Hash: hash}
	chr.Nodes = append(chr.Nodes, node)
	sort.Slice(chr.Nodes, func(i, j int) bool {
		return chr.Nodes[i].Hash < chr.Nodes[j].Hash
	})
}

// RemoveNode 删除一个缓存节点
func (chr *ConsistentHashRing) RemoveNode(ip string) {
	hash := hash(ip)
	for i, node := range chr.Nodes {
		if node.Hash == hash {
			chr.Nodes = append(chr.Nodes[:i], chr.Nodes[i+1:]...)
			break
		}
	}
}

// GetNode 查找数据键对应的缓存节点
func (chr *ConsistentHashRing) GetNode(key string) string {
	hash := hash(key)
	i := sort.Search(len(chr.Nodes), func(i int) bool {
		return chr.Nodes[i].Hash >= hash
	})
	if i == len(chr.Nodes) {
		i = 0
	}
	return chr.Nodes[i].IP
}

// hash 计算字符串的哈希值
func hash(s string) uint32 {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return (uint32(bs[0]) << 24) | (uint32(bs[1]) << 16) | (uint32(bs[2]) << 8) | uint32(bs[3])
}

func main() {
	chr := NewConsistentHashRing()

	// 添加缓存节点
	chr.AddNode("192.168.0.1")
	chr.AddNode("192.168.0.2")
	chr.AddNode("192.168.0.3")

	// 查找数据键对应的缓存节点
	keys := []string{"key1", "key2", "key3"}
	for _, key := range keys {
		node := chr.GetNode(key)
		fmt.Printf("Key: %s -> Node: %s\n", key, node)
	}

	// 删除缓存节点
	chr.RemoveNode("192.168.0.2")

	// 查找数据键对应的缓存节点
	for _, key := range keys {
		node := chr.GetNode(key)
		fmt.Printf("Key: %s -> Node: %s\n", key, node)
	}
}
