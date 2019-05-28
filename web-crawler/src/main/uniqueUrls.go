package main

import "sync"

type UniqueUrls struct {
	list map[string]bool
	mux  sync.Mutex
}

func (u *UniqueUrls) set(url string) {
	if u.list[url] {
		return
	}

	u.mux.Lock()
	defer u.mux.Unlock()
	u.list[url] = true
}
