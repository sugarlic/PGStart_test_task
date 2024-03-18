package postgre

import (
	"database/sql"

	_ "github.com/test/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}
