package repositories

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/quantum-bridge/core/pkg/squirrelizer"
)

// transactionsHistoryTable is the name of the transactions history table in the database.
const transactionsHistoryTable = "transactions_history"

// TransactionsHistoryRepository is the interface for the transactions history repository which is used to interact with the transactions history table in the database.
type TransactionsHistoryRepository interface {
	// New creates a new transactions history repository instance.
	New() TransactionsHistoryRepository
	// Get returns the first transactions history that matches the given filters.
	Get() (*TransactionsHistory, error)
	// Insert inserts the given transactions history into the database.
	Insert(transactionsHistory *TransactionsHistory) error
	// Select returns the transactions history that match the given filters.
	Select() ([]TransactionsHistory, error)
}

var (
	// transactionsHistorySelect is the select statement for the transactions history table.
	transactionsHistorySelect = sq.Select("*").From(transactionsHistoryTable)
)

// transactionsHistory is the structure that holds the transactions history repository data and filters.
type transactionsHistory struct {
	db  *squirrelizer.DB
	sql sq.SelectBuilder
}

// NewTransactionsHistory creates a new transactions history repository instance.
func NewTransactionsHistory(db *squirrelizer.DB) TransactionsHistoryRepository {
	return &transactionsHistory{
		db:  db,
		sql: transactionsHistorySelect,
	}
}

// New creates a new transactions history repository instance.
func (q *transactionsHistory) New() TransactionsHistoryRepository {
	return NewTransactionsHistory(q.db)
}

// Get returns the first transactions history that matches the given filters.
func (q *transactionsHistory) Get() (*TransactionsHistory, error) {
	var result TransactionsHistory
	err := q.db.Get(&result, q.sql)

	return &result, err
}

// Insert inserts the given transactions history into the database.
func (q *transactionsHistory) Insert(transactionsHistory *TransactionsHistory) error {
	mapping := structs.Map(transactionsHistory)

	statement := sq.Insert(transactionsHistoryTable).SetMap(mapping)

	err := q.db.Exec(statement)

	return err
}

// Select returns the transactions history that match the given filters.
func (q *transactionsHistory) Select() ([]TransactionsHistory, error) {
	var result []TransactionsHistory
	err := q.db.Select(&result, q.sql)

	return result, err
}
