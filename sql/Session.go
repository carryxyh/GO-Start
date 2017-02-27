package main

import (
	"fmt"
	"sync"
)

/*
SessionInit函数实现Session的初始化，操作成功则返回此新的Session变量
SessionRead函数返回sid所代表的Session变量，如果不存在，那么将以sid为参数调用SessionInit函数创建并返回一个新的Session变量
SessionDestroy函数用来销毁sid对应的Session变量
SessionGC根据maxLifeTime来删除过期的数据
*/

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

type Manager struct {
	cookieName  string     //private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func newManager(providerName, cookieName string, maxlifetime int64) (*Manager, err) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

func main() {
	var globalSessions *Manager
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}
