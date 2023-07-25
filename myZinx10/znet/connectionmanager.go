package znet

import (
	"errors"
	"fmt"
	"myZinx10/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	c := &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
	return c
}

// 添加连接
func (conn *ConnManager) Add(connection ziface.IConnection) {
	conn.connLock.Lock()
	defer conn.connLock.Unlock()
	conn.connections[connection.GetConnID()] = connection
	fmt.Println("[add conn to connmanager]")
}

// 删除连接
func (conn *ConnManager) Remove(connection ziface.IConnection) {
	conn.connLock.Lock()
	defer conn.connLock.Unlock()
	delete(conn.connections, connection.GetConnID())
	fmt.Println("[delete conn from connmanager]")
}

// 根据connId获取连接
func (conn *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	conn.connLock.RLock()
	defer conn.connLock.RUnlock()
	connection, ok := conn.connections[connId]
	if ok {
		return connection, nil
	}
	return nil, errors.New("[not found connection from connmanager]")

}

// 得到总数
func (conn *ConnManager) Len() int {
	fmt.Println(conn.connections)
	return len(conn.connections)
}

// 清除所有连接
func (conn *ConnManager) ClearConn() {
	conn.connLock.Lock()
	defer conn.connLock.Unlock()
	for connId, connection := range conn.connections {
		connection.Stop()
		delete(conn.connections, connId)
	}
	fmt.Println("[connmanager clear all connection]")
}
