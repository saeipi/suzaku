package client

import (
	"strconv"
	"sync"
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

func (m *Manager) Run(userIDs []string,groupId string) {
	go m.listener()
	go m.batchCreate(userIDs)
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
	var (
		client *Client
	)
	for {
		select {
		case client = <-m.unregister:
			m.unregisterClient(client)
		}
	}
}

func (m *Manager) batchCreate(userIDs []string) {
	var (
		i int
	)
	for i = 0; i < len(userIDs); i = i + 2 {
		m.newConnection(strconv.Itoa(i), strconv.Itoa(i+1))
	}
}

func (m *Manager) newConnection(uid1 string, uid2 string) {
	var (
		client1 *Client
		client2 *Client
	)
	client1 = NewClient(uid1, m)
	client2 = NewClient(uid2, m)

	m.rwLock.Lock()
	if client1.conn != nil && client2.conn != nil {
		m.clients[uid1] = client1
		m.clients[uid2] = client2
	}
	m.rwLock.Unlock()

	client1.SendUser(uid2)
}
