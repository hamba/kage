package reporter

import (
	"testing"

	"github.com/hamba/logger"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	r := &InfluxReporter{}

	Database("test")(r)

	assert.Equal(t, r.database, "test")
}

func TestMetric(t *testing.T) {
	r := &InfluxReporter{}

	Metric("kafka")(r)

	assert.Equal(t, r.metric, "kafka")
}

func TestPolicy(t *testing.T) {
	r := &InfluxReporter{}

	Policy("foobar")(r)

	assert.Equal(t, r.policy, "foobar")
}

func TestTags(t *testing.T) {
	r := &InfluxReporter{}

	Tags([]string{"foo", "bar"})(r)

	assert.Equal(t, r.tags[1], "bar")
}

func TestLog(t *testing.T) {
	log := logger.New(nil)
	r := &InfluxReporter{}

	Log(log)(r)

	assert.Equal(t, r.log, log)
}
