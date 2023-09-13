package models

import "sync"

type ServerStores struct {
	StoreHandlers map[string]StoreHandler
	StoreList     []string
	Mux           sync.RWMutex
}
