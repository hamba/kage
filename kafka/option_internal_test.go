package kafka

import (
	"testing"

	"github.com/hamba/logger"
	"github.com/stretchr/testify/assert"
)

func TestBrokers(t *testing.T) {
	brokers := []string{"127.0.0.1"}
	c := &Monitor{}

	Brokers(brokers)(c)

	assert.Equal(t, brokers, c.brokers)
}

func TestIgnoreGroups(t *testing.T) {
	i := []string{"test"}
	c := &Monitor{}

	IgnoreGroups(i)(c)

	assert.Equal(t, i, c.ignoreGroups)
}

func TestIgnoreTopics(t *testing.T) {
	i := []string{"test"}
	c := &Monitor{}

	IgnoreTopics(i)(c)

	assert.Equal(t, i, c.ignoreTopics)
}

func TestLog(t *testing.T) {
	log := logger.New(nil)
	c := &Monitor{}

	Log(log)(c)

	assert.Equal(t, log, c.log)
}

func TestOffsetChannel(t *testing.T) {
	ch := make(chan interface{})
	c := &Monitor{}

	StateChannel(ch)(c)

	assert.Equal(t, ch, c.stateCh)
}
