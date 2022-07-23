package etcd

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// NodeInfo kv中的 v 存具体结点信息
type NodeInfo struct {
	IpAddr  string `json:"ip_addr"`
	Port    int    `json:"port"`
	SvrID   string `json:"svr_id"`
	SvrName string `json:"svr_name"`
}

// RandomNum 生成随机数，用来模拟随机算法做服务发现时随机挑选一个服务实例
func RandomNum(s int64, e int64) int64 {
	//随机数如果 Seed不变 则生成的随机数一直不变
	rand.Seed(time.Now().UnixNano())
	r := rand.Int63n(e - s)
	return s + r
}


func getRegKeyPrefix(svrName string) string {
	return fmt.Sprintf("%s/%s/", "services", svrName)
}

func getRegKey(svrName string, svrID string) string {
	return fmt.Sprintf("%s%s", getRegKeyPrefix(svrName), svrID)
}

// LocalNodeCache 服务发现时将使用的结点信息缓存到本地，这里有一点不好就是如果调用过一次就会一直存在本地缓存了，需要设计一条缓存淘汰机制
type LocalNodeCache struct {
	sync.RWMutex
	// <serviceName,nodes>
	nodes map[string][]*NodeInfo
}

// 查询本地缓存是否有某个服务
func (n *LocalNodeCache) hasSvr(svrName string) bool {
	_, exist := n.nodes[svrName]
	return exist
}

// AddNode 向本地缓存中添加结点信息
func (n *LocalNodeCache) AddNode(node *NodeInfo) {
	if node == nil {
		return
	}
	n.Lock()
	defer n.Unlock()
	if !n.hasSvr(node.SvrName) {
		n.nodes[node.SvrName] = make([]*NodeInfo, 0, 3)
		n.nodes[node.SvrName] = append(n.nodes[node.SvrName], node)
		return
	}
	for idx, oldNode := range n.nodes[node.SvrName] {
		if oldNode.SvrID == node.SvrID {
			n.nodes[node.SvrName][idx] = node
			return
		}
	}
	n.nodes[node.SvrName] = append(n.nodes[node.SvrName], node)
}

// DelNode 本地缓存删除服务实例信息
func (n *LocalNodeCache) DelNode(svrName string, svrID string) {
	n.Lock()
	defer n.Unlock()
	if !n.hasSvr(svrName) {
		return
	}
	for idx, node := range n.nodes[svrName] {
		if node.SvrID == svrID {
			n.nodes[svrName] = append(n.nodes[svrName][:idx], n.nodes[svrName][idx+1:]...)
			return
		}
	}
}

// GetNode 获取服务实例，一个服务有多个实例，随机挑选一个
func (n *LocalNodeCache) GetNode(svrName string) *NodeInfo {
	nodes, exist := n.nodes[svrName]
	if !exist || len(nodes) == 0 {
		return nil
	}
	log.Printf("LocalNodeCache svrName:%v nodes len:%+v \n", svrName, len(nodes))
	pickIdx := RandomNum(0, int64(len(nodes)))

	log.Printf("LocalNodeCache pick index:%v\n", pickIdx)
	return nodes[pickIdx]
}

