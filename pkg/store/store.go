package store

import (
	"github.com/Alfagov/CesarCrypt/models"
	"github.com/Alfagov/CesarCrypt/protoStructs"
	"github.com/oklog/ulid/v2"
)

var BackendConstructors = map[string]func(string, map[string]string) models.StoreHandler{
	"fs": NewFSStoreBackend,
}

func CreateStoreConfig(name string, nShards int) *protoStructs.Config {
	return &protoStructs.Config{
		Name:           name,
		Shards:         int32(nShards),
		Version:        0,
		KeyringVersion: 0,
		Uid:            ulid.Make().String(),
	}
}
