package client

import (
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

var sendCount = 0

func (m *Manager) Run(userIDs []string, groupId string) {
	go m.listener()
	m.batchCreate(userIDs)
	for _, c := range m.clients {
		sendCount++
		if sendCount > 100 {
			break
		}
		c.SendGroup(groupId)
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
		userId string
	)
	for _, userId = range userIDs {
		m.newConnection(userId)
	}
}

func (m *Manager) newConnection(userId string) {
	var (
		client *Client
	)
	client = NewClient(userId, m)

	m.rwLock.Lock()
	if client.conn != nil {
		m.clients[userId] = client
	}
	m.rwLock.Unlock()
}
