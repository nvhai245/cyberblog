package connection

import (
    "github.com/jmoiron/sqlx"
)

var (
	// DB to use globally
	DB *sqlx.DB
)