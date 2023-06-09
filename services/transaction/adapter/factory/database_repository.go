package factory

import (
	"database/sql"

	repo "github.com/ksd-mx/mx-bank/services/transaction/adapter/repository"
	"github.com/ksd-mx/mx-bank/services/transaction/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (f *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(f.DB)
}
