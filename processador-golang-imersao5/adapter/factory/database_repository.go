package factory

import (
	"database/sql"

	repo "github.com/sidartaoss/imersao5-gateway-pagamento/adapter/repository"
	"github.com/sidartaoss/imersao5-gateway-pagamento/domain/repository"
)

type RepositoryDatabaseFactory struct {
	*sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(r.DB)	
}
