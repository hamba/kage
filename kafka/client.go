package kafka

import (
	"sync"

	"github.com/Shopify/sarama"
)

type Broker struct {
	ID        string
	Rack      string
	Connected bool
}

// Client is a kafka client.
type Client struct {
	client sarama.Client

	mu      sync.RWMutex
	brokers []Broker
	topic   map[string]int
}

// NewClient returns a kafka client with brokers.
func NewClient(addrs []string) (*Client, error) {


	c := &Client{}



	return c, nil
}

func (c *Client) refreshMetadata() error {

}

// GetBrokers returns the known brokers and their states.
func (c *Client) GetBrokers() []Broker {

}

func (c *Client) GetTopics() map[string][]TopicPartition {

}

func (c *Client) GetGroups() map[string]map[string][]GroupPartition {

}
