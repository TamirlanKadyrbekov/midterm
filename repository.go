package main

import (
	"sync"
)

var (
	mu    sync.Mutex
	Mymap = make(map[string]string)
)

func init() {
	RepositoryCreate(MyMap{"190103045", "+77779777109"})
	RepositoryCreate(MyMap{"194638902", "+77779777109"})
	RepositoryCreate(MyMap{"190745763", "+77727731028"})
	RepositoryCreate(MyMap{"193534641", "+77077076028"})
	RepositoryCreate(MyMap{"190643462", "+77707775626"})
	RepositoryCreate(MyMap{"192464570", "+77077274022"})
	RepositoryCreate(MyMap{"200303045", "+77021273072"})
	RepositoryCreate(MyMap{"190103045", "+77704524563"})

}

func RepositoryCreate(p MyMap) MyMap {
	mu.Lock()
	defer mu.Unlock()
	Mymap[p.Key] = p.Value
	return p
}

func RepositoryGet(key string) MyMap {
	mu.Lock()
	defer mu.Unlock()
	if value, ok := Mymap[key]; ok {
		return MyMap{key, value}
	}
	return MyMap{}
}

func RepositoryPut(p MyMap) MyMap {
	mu.Lock()
	defer mu.Unlock()
	Mymap[p.Key] = p.Value
	return MyMap{p.Key, p.Value}
}
