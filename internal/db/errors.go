package db

import "errors"


var (
	ErrNameDuplicate = errors.New("name duplicate")
	ErrNotFound 	= errors.New("not found")
)