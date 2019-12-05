package main

import (
	"log"
	"os"

	"github.com/hamba/cmd"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

import _ "github.com/joho/godotenv/autoload"

// Flag constants declared for CLI use.
const (
	FlagConfig = "config"

	FlagKafkaBrokers      = "kafka.brokers"
	FlagKafkaIgnoreTopics = "kafka.ignore-topics"
	FlagKafkaIgnoreGroups = "kafka.ignore-groups"

	FlagReporters = "reporters"

	FlagInflux       = "influx"
	FlagInfluxMetric = "influx.metric"
	FlagInfluxPolicy = "influx.policy"
	FlagInfluxTags   = "influx.tags"

	FlagServer = "server"
)

var version = "¯\\_(ツ)_/¯"

var agentCommand = &cli.Command{
	Name:  "agent",
	Usage: "Run the kage agent",
	Flags: cmd.Flags{
		&cli.StringFlag{
			Name:    FlagConfig,
			Usage:   "The YAML configuration file to configure ",
		},

		&cli.StringSliceFlag{
			Name:    FlagKafkaBrokers,
			Usage:   "Specify the Kafka seed brokers",
			EnvVars: []string{"KAGE_KAFKA_BROKERS"},
		},
		&cli.StringSliceFlag{
			Name:    FlagKafkaIgnoreTopics,
			Usage:   "Specify the Kafka topic patterns to ignore (may contain wildcards)",
			EnvVars: []string{"KAGE_KAFKA_IGNORE_TOPICS"},
		},
		&cli.StringSliceFlag{
			Name:    FlagKafkaIgnoreGroups,
			Usage:   "Specify the Kafka group patterns to ignore (may contain wildcards)",
			EnvVars: []string{"KAGE_KAFKA_IGNORE_GROUPS"},
		},

		&cli.StringSliceFlag{
			Name:    FlagReporters,
			Value:   cli.NewStringSlice("stdout"),
			Usage:   `"Specify the reporters to use (options: "influx", "stdout")"`,
			EnvVars: []string{"KAGE_REPORTERS"},
		},

		&cli.StringFlag{
			Name:    FlagInflux,
			Usage:   `"Specify the InfluxDB DSN (e.g. "http://user:pass@ip:port/database")"`,
			EnvVars: []string{"KAGE_INFLUX"},
		},
		&cli.StringFlag{
			Name:    FlagInfluxMetric,
			Value:   "kafka",
			Usage:   "Specify the InfluxDB metric name",
			EnvVars: []string{"KAGE_INFLUX_METRIC"},
		},
		&cli.StringFlag{
			Name:    FlagInfluxPolicy,
			Usage:   "Specify the InfluxDB metric policy",
			EnvVars: []string{"KAGE_INFLUX_POLICY"},
		},
		&cli.StringSliceFlag{
			Name:    FlagInfluxTags,
			Usage:   `"Specify additions tags to add to all metrics (e.g. "tag1=value")"`,
			EnvVars: []string{"KAGE_INFLUX_TAGS"},
		},

		&cli.BoolFlag{
			Name:    FlagServer,
			Usage:   "Start the http server",
			EnvVars: []string{"KAGE_SERVER"},
		},
	}.Merge(cmd.LogFlags, cmd.ServerFlags),
	Action: runServer,
}

func main() {
	agentCommand.Before = altsrc.InitInputSourceWithContext(agentCommand.Flags, altsrc.NewYamlSourceFromFlagFunc(FlagConfig))

	app := cli.NewApp()
	app.Name = "kage"
	app.Usage = "A Kafka monitoring agent"
	app.Version = version
	app.Commands = []*cli.Command{agentCommand}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
