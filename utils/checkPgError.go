package utils

import "github.com/jackc/pgconn"

// ErrorCodePG ...
func ErrorCodePG(err error) string {
	pgerr, ok := err.(*pgconn.PgError)
	if !ok {
		return ""
	}
	return pgerr.Code
}
