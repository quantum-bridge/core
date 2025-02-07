package repositories

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/quantum-bridge/core/pkg/squirrelizer"
)

// withdrawalsHistoryTable is the name of the withdrawals history table in the database.
const withdrawalsHistoryTable = "withdrawals_history"

// WithdrawalsHistoryRepository is the interface for the withdrawals history repository which is used to interact with the withdrawals history table in the database.
type WithdrawalsHistoryRepository interface {
	// New creates a new withdrawals history repository instance.
	New() WithdrawalsHistoryRepository
	// Get returns the first withdrawals history that matches the given filters.
	Get() (*WithdrawalsHistory, error)
	// Select returns the withdrawals history that match the given filters.
	Select() ([]WithdrawalsHistory, error)
	// Insert inserts the given withdrawals history into the database.
	Insert(withdrawalsHistory *WithdrawalsHistory) error
	// Where sets the where clause for the withdrawals history query.
	Where(column string, value interface{}) WithdrawalsHistoryRepository
	// OrderBy sets the order by for the withdrawals history query.
	OrderBy(column, direction string) WithdrawalsHistoryRepository
	// Limit sets the limit for the withdrawals history query.
	Limit(limit uint64) WithdrawalsHistoryRepository
}

var (
	// withdrawalsHistorySelect is the select statement for the withdrawals history table.
	withdrawalsHistorySelect = sq.Select("*").From(withdrawalsHistoryTable)
)

// withdrawalsHistory is the structure that holds the withdrawals history repository data and filters.
type withdrawalsHistory struct {
	db  *squirrelizer.DB
	sql sq.SelectBuilder
}

// NewWithdrawalsHistory creates a new withdrawals history repository instance.
func NewWithdrawalsHistory(db *squirrelizer.DB) WithdrawalsHistoryRepository {
	return &withdrawalsHistory{
		db:  db,
		sql: withdrawalsHistorySelect,
	}
}

// New creates a new withdrawals history repository instance.
func (q *withdrawalsHistory) New() WithdrawalsHistoryRepository {
	return NewWithdrawalsHistory(q.db)
}

// Get returns the first withdrawals history that matches the given filters.
func (q *withdrawalsHistory) Get() (*WithdrawalsHistory, error) {
	var result WithdrawalsHistory
	err := q.db.Get(&result, q.sql)

	return &result, err
}

// Select returns the withdrawals history that match the given filters.
func (q *withdrawalsHistory) Select() ([]WithdrawalsHistory, error) {
	var result []WithdrawalsHistory
	err := q.db.Select(&result, q.sql)

	return result, err
}

// Insert inserts the given withdrawals history into the database.
func (q *withdrawalsHistory) Insert(withdrawalsHistory *WithdrawalsHistory) error {
	mapping := structs.Map(withdrawalsHistory)

	statement := sq.Insert(withdrawalsHistoryTable).SetMap(mapping)

	err := q.db.Exec(statement)

	return err
}

// Where sets the where clause for the withdrawals history query.
func (q *withdrawalsHistory) Where(column string, value interface{}) WithdrawalsHistoryRepository {
	// Handle special cases for block number comparison
	if column == "block_number >=" {
		q.sql = q.sql.Where(fmt.Sprintf("block_number >= %s", value))
		return q
	}
	if column == "block_number <=" {
		q.sql = q.sql.Where(fmt.Sprintf("block_number <= %s", value))
		return q
	}

	// Default case using Eq for exact matches
	q.sql = q.sql.Where(sq.Eq{column: value})
	return q
}

// OrderBy sets the order by for the withdrawals history query.
func (q *withdrawalsHistory) OrderBy(column, direction string) WithdrawalsHistoryRepository {
	q.sql = q.sql.OrderBy(fmt.Sprintf("%s %s", column, direction))

	return q
}

// Limit sets the limit for the withdrawals history query.
func (q *withdrawalsHistory) Limit(limit uint64) WithdrawalsHistoryRepository {
	q.sql = q.sql.Limit(limit)

	return q
}
