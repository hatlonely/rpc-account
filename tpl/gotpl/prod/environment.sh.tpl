#!/usr/bin/env bash

export NAMESPACE="prod"
export NAME="rpc-account"
export REGISTRY_SERVER="{{.registry.server}}"
export REGISTRY_USERNAME="{{.registry.username}}"
export REGISTRY_PASSWORD="{{.registry.password}}"
export REGISTRY_NAMESPACE="{{.registry.namespace}}"
export MYSQL_SERVER="{{.mysql.server}}"
export MYSQL_ROOT_PASSWORD="{{.mysql.rootPassword}}"
export MYSQL_USERNAME="{{.mysql.username}}"
export MYSQL_PASSWORD="{{.mysql.password}}"
export MYSQL_DATABASE="account"
export REDIS_ADDRESS="{{.redis.addr}}"
export REDIS_PASSWORD="{{.redis.password}}"
export EMAIL_SERVER="{{.email.server}}"
export EMAIL_PASSWORD="{{.email.password}}"
export EMAIL_FROM="{{.email.from}}"
export EMAIL_PORT="{{.email.port}}"
export ELASTICSEARCH_SERVER="elasticsearch-master:9200"
export IMAGE_PULL_SECRET="hatlonely-pull-secrets"
export IMAGE_REPOSITORY="rpc-account"
export IMAGE_TAG="$(cd .. && git describe --tags | awk '{print(substr($0,2,length($0)))}')"
export REPLICA_COUNT=3
export INGRESS_HOST="k8s.account.hatlonely.com"
export INGRESS_SECRET="k8s-secret"
export K8S_CONTEXT="homek8s"
