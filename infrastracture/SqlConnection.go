package infrastracture

import "database/sql"

type SqlConnection struct {
	Db *sql.DB
}
