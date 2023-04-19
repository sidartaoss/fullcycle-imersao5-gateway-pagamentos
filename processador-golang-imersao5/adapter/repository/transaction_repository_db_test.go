package repository

import (
	"os"
	"testing"

	"github.com/sidartaoss/imersao5-gateway-pagamento/adapter/repository/fixture"
	"github.com/sidartaoss/imersao5-gateway-pagamento/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	// arrange
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	id := "1"
	accountId := "1"
	amount := 12.99
	status := entity.APPROVED
	errorMessage := ""

	// act
	err := repository.Insert(id, accountId, amount, status, errorMessage)

	// assert
	assert.Nil(t, err)

}
