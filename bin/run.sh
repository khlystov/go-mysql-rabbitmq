#!/usr/bin/env bash

docker run -d --restart always --name go-mysql-rabbitmq -v $PWD/river.toml:/app/etc/river.toml -v $PWD/var:/app/var go-mysql-rabbitmq