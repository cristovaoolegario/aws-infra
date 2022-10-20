# Kafka docker compose

## To run it

```shell
docker-compose up -d
```

## To access kafka container

```shell
docker exec -it KAFKA_CONTAINER_NAME bash
```

## To consume messages on topic

```shell
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest
```

## To send messages to topic

```shell
kafka-console-producer --bootstrap-server=localhost:9092 --topic=readtest
```
