# Local Kafka / Confluent Platform setup

Notes taken from [confluent community quick start](https://docs.confluent.io/platform/current/quickstart/cos-docker-quickstart.html) amoungst other places...

```bash
# get docker-compose file
curl --silent --output docker-compose.yml https://raw.githubusercontent.com/confluentinc/cp-all-in-one/6.1.1-post/cp-all-in-one-community/docker-compose.yml

# start up
docker-compose up -d

# connect to broker
docker-compose exec broker /bin/bash

# produce
kafka-console-producer --bootstrap-server localhost:9092 --topic testtopic1

# consume
kafka-console-consumer --bootstrap-server localhost:9092 --topic testtopic1 --from-beginning

# create topic
kafka-topics --bootstrap-server localhost:9092 --create --topic testtopic2
Created topic testtopic2.

# delete topic
kafka-topics --bootstrap-server localhost:9092 --delete --topic testtopic2

# list topics
kafka-topics --bootstrap-server localhost:9092 --list
__consumer_offsets
__transaction_state
_confluent-ksql-default__command_topic
_schemas
default_ksql_processing_log
docker-connect-configs
docker-connect-offsets
docker-connect-status
testtopic1
testtopic2
```

## rest api examples

see: 
<https://docs.confluent.io/platform/current/kafka-rest/api.html> for more.

```bash
curl http://localhost:8082/v3/clusters/VA0fX_xnRiSvDvpr_8vMpw/topics
```
