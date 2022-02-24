package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPTTime(t *testing.T) {
	p, err := NewPTTime("test", "...test...")
	if assert.NoError(t,err){
		r,err:=p.Sign(context.Background())
		if assert.NoError(t,err){
			ck:=p.CheckSuccess(context.Background(),r)
			assert.Equal(t,CheckStateUnknown,ck)
		}
	}
}
