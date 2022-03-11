package kafka

import (
	"api/config/setting"
	"api/util/log"
	"encoding/json"
	"github.com/Shopify/sarama"
)



type producerStruct struct {
	Producer sarama.AsyncProducer
}

type consumerStruct struct {
	Consumer sarama.Consumer
}
var ProducerStruct producerStruct
var Consumer consumerStruct

func OpenConnect() {

	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, _ := sarama.NewAsyncProducer([]string{setting.ConfigKafka.KafkaIp}, config)
	ProducerStruct.Producer = producer

	consumer, _ := sarama.NewConsumer([]string{setting.ConfigKafka.KafkaIp}, config)
	Consumer.Consumer = consumer

}

func Push(topic string, data interface{}) error {
	var mesg *sarama.ProducerMessage
	var producer = ProducerStruct.Producer

	if jsonData, err := json.Marshal(data); err != nil {
		return err
	} else {
		msg := &sarama.ProducerMessage{
			Topic: topic,
		}
		msg.Value = sarama.ByteEncoder(jsonData)
		msg.Key = sarama.ByteEncoder(topic)
		producer.Input() <- msg
		isSuccess := true

		func(p sarama.AsyncProducer) {
			select {
			case mesg = <-p.Successes():

			case fail := <-p.Errors():
				isSuccess = false
				err = fail.Err
			}
		}(producer)

		if !isSuccess {
			return err
		}
	}
	return nil
}

func Listen(topic string, handler func(value []byte) (isBreak bool)) {
	var (
		pc    sarama.PartitionConsumer
		value []byte
	)

	var  consumer = Consumer.Consumer
	if partitionList, err := consumer.Partitions(topic); err != nil {
		log.Error(err)
	} else {
		for partition := range partitionList {
			if pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest); err != nil {
				log.Error(err)
			} else {
				for {
					select {
					case msg := <-pc.Messages():
						value = msg.Value
						break
					}
					if handler(value) {
						break
					}
				}
			}
		}
	}
}




