package client

import (
	"strconv"
	"sync"
	"time"
)

type Manager struct {
	rwLock     sync.RWMutex
	unregister chan *Client
	clients    map[string]*Client
}

func NewManager() (mgr *Manager) {
	mgr = &Manager{clients: make(map[string]*Client), unregister: make(chan *Client, 100)}
	return
}

func (m *Manager) Run() {
	var (
		i int
	)
	go m.listener()

	for i = 0; i < 10000; i = i + 2 {
		go m.newConnection(strconv.Itoa(i), strconv.Itoa(i+1))
	}
}

func (m *Manager) unregisterClient(client *Client) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	var (
		ok bool
	)
	if _, ok = m.clients[client.userID]; ok {
		delete(m.clients, client.userID)
	}
}

func (m *Manager) listener() {
	ticker := time.NewTicker(30 * time.Minute)
	var (
		client *Client
	)
	for {
		select {
		case client = <-m.unregister:
			m.unregisterClient(client)
		case <-ticker.C:
			m.batchCreate(2000)
		}
	}
}

func (m *Manager) batchCreate(count int) {
	var (
		i int
	)
	for i = 0; i < count; i = i + 2 {
		go m.newConnection(strconv.Itoa(i), strconv.Itoa(i+1))
	}
}

func (m *Manager) newConnection(uid1 string, uid2 string) {
	var (
		client1 *Client
		client2 *Client
	)
	uid1 = "uid" + uid1
	uid2 = "uid" + uid2
	client1 = NewClient(uid1, m)
	client2 = NewClient(uid2, m)

	m.rwLock.Lock()
	m.clients[uid1] = client1
	m.clients[uid2] = client2
	m.rwLock.Unlock()

	client1.SendUser(uid2)
}
