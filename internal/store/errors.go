package store

import "fmt"

var (
	ErrRecordNotFound      = fmt.Errorf("no such record")
	ErrRollbackTransaction = fmt.Errorf("unable to rollback transaction")
)
