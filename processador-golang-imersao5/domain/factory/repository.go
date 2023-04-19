package factory

import "github.com/sidartaoss/imersao5-gateway-pagamento/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
