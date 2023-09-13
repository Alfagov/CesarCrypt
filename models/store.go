package models

import "github.com/Alfagov/CesarCrypt/protoStructs"

type SecretStore struct {
	Name    string
	Type    string
	Secrets []*Secret
}

type CreateStoreBackendRequest struct {
	Name           string
	RootEncryption string `json:"rootEncryption"`
	Shamir         bool   `json:"shamir"`
	Settings       map[string]string
}

type KeyRing struct {
	SecretKey map[string]string
}

type UnlockStoreRequest struct {
	Shards          []string
	StoreAttributes map[string]string
}

type StoreHandler interface {
	Unlock(key []byte) error
	CreateStore(config *protoStructs.Config, keyring *protoStructs.KeyRing) error
	Lock() error
	ReadSecret() error
	WriteSecret() error
	ReadConfig() (*protoStructs.Config, error)
	IsOpen() bool
}
