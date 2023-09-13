package store

import (
	"fmt"
	"github.com/Alfagov/CesarCrypt/models"
	"github.com/Alfagov/CesarCrypt/pkg/encryption"
	"github.com/Alfagov/CesarCrypt/protoStructs"
	"google.golang.org/protobuf/proto"
	"os"
	"sync"
)

type FSSecretBackend struct {
	Name        string
	KeyringPath string
	ConfigPath  string
	Path        string
	Secrets     []*models.Secret
	Unsealed    bool
	Mux         sync.RWMutex
}

func NewFSStoreBackend(name string, attrs map[string]string) models.StoreHandler {
	return &FSSecretBackend{
		Name:        name,
		Path:        attrs["path"],
		ConfigPath:  attrs["path"] + fmt.Sprintf("/configs/config-%s.json", attrs["name"]),
		KeyringPath: attrs["path"] + fmt.Sprintf("/keys/keyRing-%s.json", attrs["name"]),
	}
}

func (f *FSSecretBackend) CreateStore(config *protoStructs.Config, keyring *protoStructs.KeyRing) error {
	f.Mux.Lock()
	defer f.Mux.Unlock()

	if _, err := os.Stat(f.Path); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(f.Path + "/configs"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(f.Path + "/keys"); os.IsNotExist(err) {
		return err
	}

	err := f.writeConfig(config)
	if err != nil {
		return err
	}

	err = f.writeKeyring(keyring, keyring.RootKey)
	if err != nil {
		return err
	}

	return nil
}

func (f *FSSecretBackend) writeKeyring(keyring *protoStructs.KeyRing, key string) error {
	keyringBytes, err := proto.Marshal(keyring)
	if err != nil {
		return err
	}

	encryptedKeyRing := encryption.Encrypt(string(keyringBytes), key)

	err = os.WriteFile(f.KeyringPath, []byte(encryptedKeyRing), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f *FSSecretBackend) writeConfig(config *protoStructs.Config) error {

	configBytes, err := proto.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(f.ConfigPath, configBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f *FSSecretBackend) ReadConfig() (*protoStructs.Config, error) {
	file, err := os.ReadFile(f.ConfigPath)
	if err != nil {
		return nil, err
	}

	var config protoStructs.Config
	err = proto.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (f *FSSecretBackend) ReadKeyring() (string, error) {
	file, err := os.ReadFile(f.KeyringPath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func (f *FSSecretBackend) IsOpen() bool {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.Unsealed
}

func (f *FSSecretBackend) Unlock(key []byte) error {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	if key == nil || len(key) == 0 {
		return nil
	}

	file, err := os.ReadFile(f.Path)
	if err != nil {
		return err
	}

	_ = encryption.Decrypt(string(file), string(key))

	f.Unsealed = true

	return nil
}

func (f *FSSecretBackend) Lock() error {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	panic("implement me")
}

func (f *FSSecretBackend) ReadSecret() error {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	panic("implement me")
}

func (f *FSSecretBackend) WriteSecret() error {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	panic("implement me")
}
