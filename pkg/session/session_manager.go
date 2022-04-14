package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Provider interface {
	//初始化一个session，id根据需要生成后传入
	InitSession(sid string, maxAge int64) (Session, error)
	//根据sid，获得当前session
	SetSession(session Session) error
	//销毁session
	DestroySession(sid string) error
	//回收
	GCSession()
}

type FromMemory struct {
	lock     sync.Mutex
	sessions map[string]Session
}

func newFromMemory() *FromMemory {
	return &FromMemory{
		sessions: make(map[string]Session, 0),
	}
}

func (fm *FromMemory) InitSession(sid string, maxAge int64) (Session, error) {
	fm.lock.Lock()
	defer fm.lock.Unlock()

	newSession := newSessionFromMemory()
	newSession.sid = sid
	if maxAge != 0 {
		newSession.maxAge = maxAge
	}
	newSession.lastAccessedTime = time.Now()

	fm.sessions[sid] = newSession
	fmt.Println(fm.sessions)
	return newSession, nil
}

func (fm *FromMemory) SetSession(session Session) error {
	fm.sessions[session.GetId()] = session
	return nil
}

func (fm *FromMemory) DestroySession(sid string) error {
	if _, ok := fm.sessions[sid]; ok {
		delete(fm.sessions, sid)
		return nil
	}
	return nil
}

func (fm *FromMemory) GCSession() {
	sessions := fm.sessions
	if len(sessions) < 1 {
		return
	}
	for k, v := range sessions {
		t := (v.(*SessionFromMemory).lastAccessedTime.Unix()) + (v.(*SessionFromMemory).maxAge)
		if t < time.Now().Unix() {
			delete(fm.sessions, k)
		}
	}
}

type SessionManager struct {
	cookieName string
	storage    Provider
	maxAge     int64
	lock       sync.Mutex
}

func NewSessionMange() *SessionManager {
	SessionManager := &SessionManager{
		cookieName: "lz_cookie",
		storage:    newFromMemory(),
		maxAge:     1800,
	}

	go SessionManager.GC()
	return SessionManager
}

func (m *SessionManager) GetCookieN() string {
	return m.cookieName
}

const COOKIE_MAX_MAX_AGE = time.Hour * 24 / time.Second // 单位：秒。
func (m *SessionManager) BeginSession(w http.ResponseWriter, r *http.Request) Session {
	m.lock.Lock()
	defer m.lock.Unlock()
	fmt.Println("cookie-name:", m.cookieName)
	cookie, err := r.Cookie(m.cookieName)
	fmt.Println(cookie, err)
	maxAge2 := int(COOKIE_MAX_MAX_AGE)
	if err != nil || cookie.Value == "" {
		fmt.Println("----------> current session not exists")
		sid := m.randomId()

		session, _ := m.storage.InitSession(sid, m.maxAge)
		maxAge := m.maxAge

		if maxAge == 0 {
			maxAge = session.(*SessionFromMemory).maxAge
		}

		//设置cookie

		//maxAge2 := int(COOKIE_MAX_MAX_AGE)
		uid_cookie := &http.Cookie{
			Name:     m.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: false,
			MaxAge:   maxAge2,
		}
		http.SetCookie(w, uid_cookie) //设置到响应中
		return session
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session := m.storage.(*FromMemory).sessions[sid]
		fmt.Println("sesssion----->", session)
		if session == nil {
			fmt.Println("-----------> current session is nil")
			//创建一个
			//sid := m.randomId()
			//根据保存session方式，如内存，数据库中创建
			newSession, _ := m.storage.InitSession(sid, m.maxAge) //该方法有自己的锁，多处调用到

			maxAge := m.maxAge

			if maxAge == 0 {
				maxAge = newSession.(*SessionFromMemory).maxAge
			}
			//用session的ID于cookie关联
			//cookie名字和失效时间由session管理器维护
			newCookie := http.Cookie{
				Name: m.cookieName,
				//这里是并发不安全的，但是这个方法已上锁
				Value:    url.QueryEscape(sid), //转义特殊符号@#￥%+*-等
				Path:     "/",
				HttpOnly: true,
				MaxAge:   maxAge2,
				Expires:  time.Now().Add(time.Duration(maxAge)),
			}
			http.SetCookie(w, &newCookie) //设置到响应中
			fmt.Println("-----------> current session exists")
			return newSession
		}
		return session
	}
}

//通过ID获取session
func (m *SessionManager) GetSessionById(sid string) Session {
	session := m.storage.(*FromMemory).sessions[sid]
	return session
}

//是否内存中存在
func (m *SessionManager) MemoryIsExists(sid string) bool {
	_, ok := m.storage.(*FromMemory).sessions[sid]
	if ok {
		return true
	}
	return false
}

//手动销毁session，同时删除cookie
func (m *SessionManager) Destroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		m.lock.Lock()
		defer m.lock.Unlock()

		sid, _ := url.QueryUnescape(cookie.Value)
		m.storage.DestroySession(sid)

		cookie2 := http.Cookie{
			MaxAge:  0,
			Name:    m.cookieName,
			Value:   "",
			Path:    "/",
			Expires: time.Now().Add(time.Duration(0)),
		}

		http.SetCookie(w, &cookie2)
	}
}

func (m *SessionManager) Update(w http.ResponseWriter, r *http.Request) {
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)
	if err != nil {
		return
	}

	t := time.Now()
	sid, _ := url.QueryUnescape(cookie.Value)

	sessions := m.storage.(*FromMemory).sessions
	session := sessions[sid].(*SessionFromMemory)
	session.lastAccessedTime = t

	if m.maxAge != 0 {
		cookie.MaxAge = int(m.maxAge)
	} else {
		cookie.MaxAge = int(session.maxAge)
	}
	http.SetCookie(w, cookie)
}

func (m *SessionManager) randomId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (m *SessionManager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.storage.GCSession()
	age2 := int(60 * time.Second)
	time.AfterFunc(time.Duration(age2), func() {
		m.GC()
	})
}
