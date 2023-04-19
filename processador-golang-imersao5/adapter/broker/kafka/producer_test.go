package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sidartaoss/imersao5-gateway-pagamento/adapter/presenter/transaction"
	"github.com/sidartaoss/imersao5-gateway-pagamento/domain/entity"
	"github.com/sidartaoss/imersao5-gateway-pagamento/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	// arrange
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "no limit for this transaction",
	}

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())

	// act
	err := producer.Publish(expectedOutput, []byte("1"), "test")

	// assert
	assert.Nil(t, err)
}
