package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/hamba/cmd"
	"github.com/hamba/pkg/log"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/msales/kage"
	"github.com/msales/kage/kafka"
	"github.com/msales/kage/reporter"
	"github.com/msales/kage/store"
	"github.com/urfave/cli/v2"
)

// Application =============================

func newApplication(c *cmd.Context) (*kage.Application, error) {
	logger := c.Logger()

	memStore, err := store.New()
	if err != nil {
		return nil, err
	}

	reporters, err := newReporters(c.Context, logger)
	if err != nil {
		return nil, err
	}

	monitor, err := kafka.New(
		kafka.Brokers(c.StringSlice(FlagKafkaBrokers)),
		kafka.IgnoreTopics(c.StringSlice(FlagKafkaIgnoreTopics)),
		kafka.IgnoreGroups(c.StringSlice(FlagKafkaIgnoreGroups)),
		kafka.StateChannel(memStore.Channel()),
		kafka.Log(logger),
	)
	if err != nil {
		return nil, err
	}

	app := kage.NewApplication()
	app.Store = memStore
	app.Reporters = reporters
	app.Monitor = monitor
	app.Logger = logger

	return app, nil
}

// Reporters ===============================

// newReporters creates reporters from the config.
func newReporters(c *cli.Context, logger log.Logger) (*kage.Reporters, error) {
	rs := &kage.Reporters{}

	for _, name := range c.StringSlice(FlagReporters) {
		switch name {
		case "influx":
			r, err := newInfluxReporter(c, logger)
			if err != nil {
				return nil, err
			}
			rs.Add(name, r)

		case "stdout":
			r := reporter.NewConsoleReporter(os.Stdout)
			rs.Add(name, r)

		default:
			return nil, fmt.Errorf("unknown reporter \"%s\"", name)
		}
	}

	return rs, nil
}

// newInfluxReporter create a new InfluxDB reporter.
func newInfluxReporter(c *cli.Context, logger log.Logger) (kage.Reporter, error) {
	dsn, err := url.Parse(c.String(FlagInflux))
	if err != nil {
		return nil, err
	}

	if dsn.User == nil {
		dsn.User = &url.Userinfo{}
	}

	addr := dsn.Scheme + "://" + dsn.Host
	username := dsn.User.Username()
	password, _ := dsn.User.Password()

	influx, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	db := strings.Trim(dsn.Path, "/")

	tags, err := cmd.SplitTags(c.StringSlice(FlagInfluxTags), "=")
	if err != nil {
		return nil, err
	}

	return reporter.NewInfluxReporter(influx,
		reporter.Database(db),
		reporter.Metric(c.String(FlagInfluxMetric)),
		reporter.Policy(c.String(FlagInfluxPolicy)),
		reporter.Tags(tags),
		reporter.Log(logger),
	), nil
}
