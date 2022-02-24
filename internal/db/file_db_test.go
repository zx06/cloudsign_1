package db

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zx06/cloudsign/internal/provider"
)

func createConfig(db *FileDB) error {
	testProvider, err := provider.NewSample(
		"test",
		&provider.SampleCfg{
			URL:    "http://test.com",
			Method: "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body:          `{"name":"test"}`,
			RespSuccess:   `{"code":0,"msg":"success"}`,
			RespFailure:   `{"code":1,"msg":"failure"}`,
			RespDuplicate: `{"code":2,"msg":"duplicate"}`,
		})
	if err != nil {
		return err
	}
	err = db.CreateConfig(context.Background(), &testProvider.SignConfig)
	if err != nil {
		return err
	}
	return nil
}

func TestFileDB_CreateConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.json")
	db, err := NewFileDB(path)
	if assert.NoError(t, err) {
		err := createConfig(db)
		assert.NoError(t, err)
	}
}

func TestFileDB_GetConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.json")
	db, err := NewFileDB(path)
	if assert.NoError(t, err) {
		_, err := db.GetConfig(context.Background(), "test")
		assert.EqualError(t, err, ErrNotFound.Error())
		err = createConfig(db)
		assert.NoError(t, err)
		cfg, err := db.GetConfig(context.Background(), "test")
		assert.NoError(t, err)
		assert.Equal(t, "sample", cfg.ProviderName)
		assert.Equal(t, "test", cfg.Name)
	}
}

func TestFileDB_UpdateConfig(t *testing.T) {

}

func TestFileDB_DeleteConfig(t *testing.T) {

}

func TestFileDB_ListConfig(t *testing.T) {

}
