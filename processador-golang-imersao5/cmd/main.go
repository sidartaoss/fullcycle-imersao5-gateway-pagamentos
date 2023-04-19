package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sidartaoss/imersao5-gateway-pagamento/adapter/broker/kafka"
	"github.com/sidartaoss/imersao5-gateway-pagamento/adapter/factory"
	"github.com/sidartaoss/imersao5-gateway-pagamento/adapter/presenter/transaction"
	"github.com/sidartaoss/imersao5-gateway-pagamento/usecase/process_transaction"
)

func main() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/"+os.Getenv("MYSQL_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.mechanisms":   os.Getenv("SASL_MECHANISMS"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
	}

	kafkaPresenter := transaction.NewTransactionKafkaPresenter()

	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	msgChan := make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.mechanisms":   os.Getenv("SASL_MECHANISMS"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
		"client.id":         "goapp",
		"group.id":          "goapp",
	}

	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		err = json.Unmarshal(msg.Value, &input)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("processing message from channel")
		_, err = usecase.Execute(input)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
