package global

import (
	"sync"
)

type SyncMap struct {
	m sync.Map
}

func (m SyncMap) SetMap(key any, value any) {
	m.m.Store(key, value)
}

func (m SyncMap) SelectMap(key any) bool {
	_, ok := m.m.Load(key)
	return ok
}

func (m SyncMap) DeleteMap(key any) {
	m.m.Delete(key)
}

var ImportProject SyncMap
