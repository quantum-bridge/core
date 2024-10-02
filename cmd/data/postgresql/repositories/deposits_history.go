package repositories

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/quantum-bridge/core/pkg/squirrelizer"
)

// depositsHistoryTable is the name of the deposits history table in the database.
const depositsHistoryTable = "deposits_history"

// DepositsHistoryRepository is the interface for the deposits history repository which is used to interact with the deposits history table in the database.
type DepositsHistoryRepository interface {
	// New creates a new deposits history repository instance.
	New() DepositsHistoryRepository
	// Get returns the first deposits history that matches the given filters.
	Get() (*DepositsHistory, error)
	// Select returns the deposits history that match the given filters.
	Select() ([]DepositsHistory, error)
	// Insert inserts the given deposits history into the database.
	Insert(depositsHistory *DepositsHistory) error
	// Where sets the where clause for the deposits history query.
	Where(column, value string) DepositsHistoryRepository
	// OrderBy sets the order by for the deposits history query.
	OrderBy(column, direction string) DepositsHistoryRepository
	// Limit sets the limit for the deposits history query.
	Limit(limit uint64) DepositsHistoryRepository
}

var (
	// depositsHistorySelect is the select statement for the deposits history table.
	depositsHistorySelect = sq.Select("*").From(depositsHistoryTable)
)

// depositsHistory is the structure that holds the deposits history repository data and filters.
type depositsHistory struct {
	db  *squirrelizer.DB
	sql sq.SelectBuilder
}

// NewDepositsHistory creates a new deposits history repository instance.
func NewDepositsHistory(db *squirrelizer.DB) DepositsHistoryRepository {
	return &depositsHistory{
		db:  db,
		sql: depositsHistorySelect,
	}
}

// New creates a new deposits history repository instance.
func (q *depositsHistory) New() DepositsHistoryRepository {
	return NewDepositsHistory(q.db)
}

// Get returns the first deposits history that matches the given filters.
func (q *depositsHistory) Get() (*DepositsHistory, error) {
	var result DepositsHistory
	err := q.db.Get(&result, q.sql)

	return &result, err
}

// Select returns the deposits history that match the given filters.
func (q *depositsHistory) Select() ([]DepositsHistory, error) {
	var result []DepositsHistory
	err := q.db.Select(&result, q.sql)

	return result, err
}

// Insert inserts the given deposits history into the database.
func (q *depositsHistory) Insert(depositsHistory *DepositsHistory) error {
	mapping := structs.Map(depositsHistory)

	statement := sq.Insert(depositsHistoryTable).SetMap(mapping)

	err := q.db.Exec(statement)

	return err
}

// Where sets the where clause for the deposits history query.
func (q *depositsHistory) Where(column, value string) DepositsHistoryRepository {
	q.sql = q.sql.Where(sq.Eq{column: value})

	return q
}

// OrderBy sets the order by for the deposits history query.
func (q *depositsHistory) OrderBy(column, direction string) DepositsHistoryRepository {
	q.sql = q.sql.OrderBy(fmt.Sprintf("%s %s", column, direction))

	return q
}

// Limit sets the limit for the deposits history query.
func (q *depositsHistory) Limit(limit uint64) DepositsHistoryRepository {
	q.sql = q.sql.Limit(limit)

	return q
}
