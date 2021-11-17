package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log_agent/config"
)

var producer sarama.SyncProducer

func init()  {
	saramaConfig:= sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaConfig.Producer.Return.Successes = true
	var err error
	producer, err = sarama.NewSyncProducer([]string{config.AppConfigObj.KafkaConfig.Address}, saramaConfig)
	if err != nil {
		fmt.Printf("连接kafka出错：%v", err)
		return
	}
}

func SendMsg(msg string, topic string) error  {
	message, i, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	})
	if err != nil {
		return err
	}
	fmt.Println("发送消息成功：", msg,message, i, topic)
	return nil
}