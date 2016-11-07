package main

import (
	"sync"
	"container/list"
	"net"
)

type Session struct {
	Conn net.Conn
	Fd   int
}

type Room struct {
	mutex sync.RWMutex
	users list.List
	index int
	size  int
}

func (this *Room)Join(conn net.Conn) (*Session) {
	this.mutex.Lock()
	session := &Session{}
	session.Conn = conn;
	session.Fd = this.index
	this.index++
	this.size++;
	this.users.PushBack(session)
	this.mutex.Unlock()
	return session
}

func (this *Room)Leave(session *Session) {
	this.mutex.Lock()
	for it := this.users.Front(); it != nil; it = it.Next() {
		if ( it.Value.(*Session).Fd == session.Fd) {
			this.users.Remove(it)
			this.size--;
			this.mutex.Unlock()
			return
		}
	}

	this.mutex.Unlock()
}

func (this *Room)Broadcast(msg string) {
	this.mutex.RLock()
	for it := this.users.Front(); it != nil; it = it.Next() {
		it.Value.(*Session).Conn.Write([]byte(msg))
	}
	this.mutex.RUnlock()
}

func (this *Room)SendToFd(fd int, msg string) {
	this.mutex.RLock()
	for it := this.users.Front(); it != nil; it = it.Next() {
		if (it.Value.(*Session).Fd == fd) {
			it.Value.(*Session).Conn.Write([]byte(msg))
			return
		}
	}
	this.mutex.RUnlock()
}