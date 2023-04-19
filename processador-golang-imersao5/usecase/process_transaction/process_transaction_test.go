package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_broker "github.com/sidartaoss/imersao5-gateway-pagamento/adapter/broker/mock"
	"github.com/sidartaoss/imersao5-gateway-pagamento/domain/entity"
	mock_repository "github.com/sidartaoss/imersao5-gateway-pagamento/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionExecuteInvalidCreditCard(t *testing.T) {
	// arrange
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "400000000000",
		CreditCardName:            "Sidarta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().AddDate(4, 0, 0).Year(),
		CreditCardCVV:             123,
		Amount:                    float64(200),
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	topic := "transactions_result"

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), topic)

	usecase := NewProcessTransaction(repositoryMock, producerMock, topic)

	// act
	output, err := usecase.Execute(input)

	// assert
	assert.NotNil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, expectedOutput, output)

}

func TestProcessTransactionExecuteRejectedTransaction(t *testing.T) {
	// arrange
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4538520392582409",
		CreditCardName:            "Sidarta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().AddDate(4, 0, 0).Year(),
		CreditCardCVV:             123,
		Amount:                    float64(1200),
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "no limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	topic := "transactions_result"

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), topic)

	usecase := NewProcessTransaction(repositoryMock, producerMock, topic)

	// act
	output, err := usecase.Execute(input)

	// assert
	assert.NotNil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, expectedOutput, output)

}

func TestProcessTransactionExecuteApprovedTransaction(t *testing.T) {
	// arrange
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4538520392582409",
		CreditCardName:            "Sidarta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().AddDate(4, 0, 0).Year(),
		CreditCardCVV:             123,
		Amount:                    float64(900),
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	topic := "transactions_result"

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), topic)

	usecase := NewProcessTransaction(repositoryMock, producerMock, topic)

	// act
	output, err := usecase.Execute(input)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, expectedOutput, output)

}
