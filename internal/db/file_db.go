package db

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/zx06/cloudsign/internal/provider"
)

var _ DB = (*FileDB)(nil)

type JsonConfig struct {
	Data       []provider.SignConfig
	LastUpdate time.Time
}

type FileDB struct {
	path string
	mu   sync.RWMutex
}

func NewFileDB(path string) (*FileDB, error) {
	return &FileDB{
		path: path,
	}, nil
}

func (db *FileDB) CreateConfig(ctx context.Context, cfg *provider.SignConfig) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	err := db.createDefault(ctx)
	if err != nil {
		return err
	}
	// read config
	var jsonCfg JsonConfig
	fr, err := os.Open(db.path)
	if err != nil {
		return err
	}
	defer fr.Close()
	err = json.NewDecoder(fr).Decode(&jsonCfg)
	if err != nil {
		return err
	}
	// check name duplicate
	for _, v := range jsonCfg.Data {
		if v.Name == cfg.Name {
			return ErrNameDuplicate
		}
	}
	// append config
	jsonCfg.Data = append(jsonCfg.Data, *cfg)
	jsonCfg.LastUpdate = time.Now()
	// write config
	fw, err := os.OpenFile(db.path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fw.Close()
	err = json.NewEncoder(fw).Encode(jsonCfg)
	if err != nil {
		return err
	}
	return nil
}

func (db *FileDB) GetConfig(ctx context.Context, name string) (*provider.SignConfig, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	err := db.createDefault(ctx)
	if err != nil {
		return nil, err
	}
	// read config
	var jsonCfg JsonConfig
	fr, err := os.Open(db.path)
	if err != nil {
		return nil, err
	}
	defer fr.Close()
	err = json.NewDecoder(fr).Decode(&jsonCfg)
	if err != nil {
		return nil, err
	}
	// check name
	for _, v := range jsonCfg.Data {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, ErrNotFound
}

func (db *FileDB) UpdateConfig(ctx context.Context, cfg *provider.SignConfig) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	err := db.createDefault(ctx)
	if err != nil {
		return err
	}
	// read config
	var jsonCfg JsonConfig
	fr, err := os.Open(db.path)
	if err != nil {
		return err
	}
	defer fr.Close()
	err = json.NewDecoder(fr).Decode(&jsonCfg)
	if err != nil {
		return err
	}
	// check name
	for i, v := range jsonCfg.Data {
		if v.Name == cfg.Name {
			jsonCfg.Data[i] = *cfg
			jsonCfg.LastUpdate = time.Now()
			// write config
			fw, err := os.OpenFile(db.path, os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				return err
			}
			defer fw.Close()
			err = json.NewEncoder(fw).Encode(jsonCfg)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNotFound
}

func (db *FileDB) DeleteConfig(ctx context.Context, name string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	err := db.createDefault(ctx)
	if err != nil {
		return err
	}
	// read config
	var jsonCfg JsonConfig
	fr, err := os.Open(db.path)
	if err != nil {
		return err
	}
	defer fr.Close()
	err = json.NewDecoder(fr).Decode(&jsonCfg)
	if err != nil {
		return err
	}
	// check name
	for i, v := range jsonCfg.Data {
		if v.Name == name {
			jsonCfg.Data = append(jsonCfg.Data[:i], jsonCfg.Data[i+1:]...)
			jsonCfg.LastUpdate = time.Now()
			// write config
			fw, err := os.OpenFile(db.path, os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				return err
			}
			defer fw.Close()
			err = json.NewEncoder(fw).Encode(jsonCfg)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNotFound
}

func (db *FileDB) ListConfig(ctx context.Context) ([]provider.SignConfig, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	err := db.createDefault(ctx)
	if err != nil {
		return nil, err
	}
	// read config
	var jsonCfg JsonConfig
	fr, err := os.Open(db.path)
	if err != nil {
		return nil, err
	}
	defer fr.Close()
	err = json.NewDecoder(fr).Decode(&jsonCfg)
	if err != nil {
		return nil, err
	}
	return jsonCfg.Data, nil
}

func (db *FileDB) createDefault(ctx context.Context) error {
	// read config
	var jsonCfg = JsonConfig{
		LastUpdate: time.Now(),
	}
	fr, err := os.Open(db.path)
	if err == nil {
		fr.Close()
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	defer fr.Close()
	fw, err := os.OpenFile(db.path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fw.Close()
	return json.NewEncoder(fw).Encode(&jsonCfg)
}
