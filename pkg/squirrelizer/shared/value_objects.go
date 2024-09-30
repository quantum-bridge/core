package shared

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Rebind replaces the bind variables in the query with the correct bind variables for the database.
func Rebind(statement string) string {
	return sqlx.Rebind(sqlx.BindType("postgres"), statement)
}

// Build builds the SQL query from the squirrel query builder.
func Build(sqlizer squirrel.Sqlizer) (string, []interface{}, error) {
	sqlStatement, args, err := sqlizer.ToSql()
	if err != nil {
		err = errors.Wrap(err, "failed to build SQL query")
	}

	return sqlStatement, args, err
}
