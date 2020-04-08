package kafkaclient

import (
	"errors"

	"fmt"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kazoo-go"
)

type KafkaConfig struct {
	*sarama.Config
	Zookeeper *kazoo.Config
}

func NewConfig() *KafkaConfig {
	config := &KafkaConfig{}

	config.Config = sarama.NewConfig()
	config.Zookeeper = kazoo.NewConfig()

	return config
}

func (conf *KafkaConfig) Validate() error {
	if conf.Zookeeper.Timeout <= 0 {
		return sarama.ConfigurationError("ZookeeperTimeout should have a duration > 0")
	}

	if conf.Config != nil {
		if err := conf.Config.Validate(); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

type HashProductGroup struct {
	producer sarama.SyncProducer
}

func NewHashProductGroup(zookeeper []string) (*HashProductGroup, error) {
	var kz *kazoo.Kazoo
	var err error
	product := &HashProductGroup{}
	config := NewConfig()

	if kz, err = kazoo.NewKazoo(zookeeper, config.Zookeeper); err != nil {
		return nil, errors.New("You need to provide at least one zookeeper node address!")
	}

	brokers, err := kz.BrokerList()
	if err != nil {
		kz.Close()
		return nil, errors.New("You need to provide at least one zookeeper node address!")
	}

	config.Config.Producer.RequiredAcks = sarama.WaitForAll
	config.Config.Producer.Return.Successes = true
	config.Config.Producer.Partitioner = sarama.NewHashPartitioner
	product.producer, err = sarama.NewSyncProducer(brokers, config.Config)
	if err != nil {
		fmt.Println("FAILED to open the producer:", err)
		return nil, err
	}

	return product, nil
}

func (product *HashProductGroup) PushMsg(topics string, key string, value []byte) error {
	keyEncoder := sarama.StringEncoder(key)
	valueEncoder := sarama.ByteEncoder(value)

	_, _, err := product.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topics,
		Key:   keyEncoder,
		Value: valueEncoder,
	})

	return err
}
