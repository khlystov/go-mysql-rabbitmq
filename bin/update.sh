#!/usr/bin/env bash

docker stop go-mysql-rabbitmq && docker rm go-mysql-rabbitmq || true
docker build -t go-mysql-rabbitmq .
docker run -d --name go-mysql-rabbitmq -v $PWD/river.toml:/app/etc/river.toml -v $PWD/var:/app/var go-mysql-rabbitmq