package session

import (
	"sync"
	"time"
)

type Session interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Remove(key interface{}) error
	GetId() string
}

type SessionFromMemory struct {
	sid              string //唯一标识
	lock             sync.Mutex
	lastAccessedTime time.Time                   //最后一次访问时间
	maxAge           int64                       //超时时间
	data             map[interface{}]interface{} //主数据
}

const DEFEALT_TIME = 1800

//实例化
func newSessionFromMemory() *SessionFromMemory {
	return &SessionFromMemory{
		data:   make(map[interface{}]interface{}),
		maxAge: DEFEALT_TIME,
	}
}

func (si *SessionFromMemory) Set(key, value interface{}) {
	si.lock.Lock()
	defer si.lock.Unlock()
	si.data[key] = value
}

func (si *SessionFromMemory) Get(key interface{}) interface{} {
	if value := si.data[key]; value != nil {
		return value
	}
	return nil
}

func (si *SessionFromMemory) Remove(key interface{}) error {
	if value := si.data[key]; value != nil {
		delete(si.data, key)
	}
	return nil
}

func (si *SessionFromMemory) GetId() string {
	return si.sid
}
