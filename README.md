![Logo](http://svg.wiersma.co.za/hamba/project?title=kage&tag=A%20Kafka%20monitoring%20agent)

[![Go Report Card](https://goreportcard.com/badge/github.com/hamba/kage)](https://goreportcard.com/report/github.com/hamba/kage)
[![Build Status](https://travis-ci.com/hamba/kage.svg?branch=master)](https://travis-ci.com/hamba/kage)
[![Coverage Status](https://coveralls.io/repos/github/hamba/kage/badge.svg?branch=master)](https://coveralls.io/github/hamba/kage?branch=master)
[![GitHub release](https://img.shields.io/github/release/hamba/kage.svg)](https://github.com/hamba/kage/releases)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hamba/kage/master/LICENSE)

## Synopsis

Kage (as in "Kafka AGEnt") reads Offset- and Lag metrics from Kafka and writes them to an InfluxDB.

## Motivation

When you're running Kafka you probably want to have monitoring as well.  
You can - of course - query the beans directly via JMX and work with that, but that requires another JVM that collects the data.  
If you're a java-shop anyway and have all that available - give it a try.
We decided that we wanted to get the metrics straight out of Kafka and feed them into InfluxDB in a configurable way - and here we are now.

## A New Beginning

In its current form, Kage as stood the test of time. It however needs to change its form from an InfluxDB
writer into a Kafka metrics reporter. To that end Kage has been moved here to start its new beginning. 
It is my hope that more community and get some love.

## Basic Installation

Grab a binary from the Releases or clone the repo and build it yourself.  
Run the binary with the configuration options found below.
```sh
kage agent <CONFIG OPTIONS>

```

## Advanced Installation

There's systemd configuration magic in examples/systemd/.  
Put the files in the appropriate directories on your machine (in case of Debian/Ubuntu that should be /lib/systemd/system 
and /lib/systemd/system-generators), remember to chmod 0755 the generator, create /etc/kage/, run `systemctl daemon-reload` 
and then you should get one service per configuration-file in /etc/kage/.

## Configuration

Kage can be configured with command line flags and environment variables. 
 
#### Command Line Flags

| Flag | Options | Multiple Allowed | Description | Environment Variable |
| ---- | ------- | ---------------- | ----------- | -------------------- |
| --log.format | logfmt, json | No | The format of logs. | LOG_FORMAT |
| --log.level | debug, info, error | No | The log level to use. | LOG_LEVEL |
| --log.tags | | Yes | A list of tags appended to every log. | LOG_TAGS |
| --kafka.brokers | | Yes | The kafka seed brokers connect to. Format: 'ip:port'. | KAGE_KAFKA_BROKERS |
| --kafka.ignore-topics | | Yes | The kafka topic patterns to ignore. This may contian wildcards. | KAGE_KAFKA_IGNORE_TOPICS |
| --kafka.ignore-groups | | Yes | The kafka consumer group patterns to ignore. This may contian wildcards. | KAGE_KAFKA_IGNORE_GROUPS |
| --reporters | influx, stdout | Yes | The reporters to use. | KAGE_REPORTERS |
| --influx | | No | The DSN of the InfluxDB server to report to. Format: http://user:pass@ip:port/database'. | KAGE_INFLUX |
| --influx.metric | | No | The measurement name to report statistics under. | KAGE_INFLUX_METRIC |
| --influx.policy | | No | The retention policy to report statistics under. | KAGE_INFLUX_POLICY |
| --influx.tags | | Yes | Additional tags to add to the statistics. Format: 'key=value' | KAGE_INFLUX_TAGS |
| --server | | No | Start the http server. | KAGE_SERVER |
| --port | | No | The port to bind to for the http server. | PORT |

##### Multi value environment variables

When using environment variables where mutltiple values are allowed, the values should be comma seperated.
E.g. `--reporters=stdout --reporters=influx` should become `KAGE_REPORTERS=stdout,influx`.

## HTTP Endpoints

Kage has an optional http server that can be enabled with the `--server` configuration. This allows health checking
as well as fetching broker and consumer group information. The endpoints are as follows:

#### GET /health

Gets the current health status of Kage. Returns a 200 status code if Kage is healthy, otherwise a 500 status code

#### GET /brokers

Get the state of all known brokers.

#### GET /brokers/health

Get the current kafka health status. Returns a 200 status code if all brokers are connected, otherwise a 500 status code
 
#### GET /topics

Get a topic offset information in json format.

#### GET /metadata

Get a topic metadata information in json format.

#### GET /consumers

Get a consumer group offset information in json format.

#### GET /consumers/:group

Get a consumer group offset information for the specified consumer group in json format, or will return with a 404 status code.

## Contributors

We're supposed to tell you how to contribute to kage here.  
Since this is github: You know the drill - open issues, fork, create PRs, ...

## TODO

 * provide ansible-templates and examples
 * set up debian packaging

## License

MIT-License. As is. No warranties whatsoever. Mileage may vary. Batteries not included.
