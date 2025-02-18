package testdb

import (
	"context"
	"database/sql"
	"testing"

	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
	"github.com/melisource/fury_go-core/pkg/sqldb/sqldbmigrate"
	"github.com/melisource/fury_go-core/pkg/sqldb/sqldbtest"
	"github.com/stretchr/testify/require"
)

func NewConn(t *testing.T) (db *sql.DB, truncate func(), teardown func()) {
	t.Helper()

	migrations, err := sqldbmigrate.GetMigrationsFS("migrations/mysql/testschema")
	require.NoError(t, err)

	schema := randstr.Chars(8)

	db, teardown = sqldbtest.Setup(
		t,
		context.Background(),
		"localhost:3306",
		migrations,
		sqldbtest.WithPassword("root"),
		sqldbtest.WithUser("root"),
		sqldbtest.WithDatabaseName(schema),
	)

	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = ?", schema)
	defer rows.Close()
	require.NoError(t, err)

	var tables []string

	for rows.Next() {
		var table string

		require.NoError(t, rows.Scan(&table))
		tables = append(tables, table)
	}

	truncate = func() {
		sqldbtest.TruncateTables(t, db, tables...)
	}

	return db, truncate, teardown
}
