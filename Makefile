BINARY_NAME = messaggio_test
IMAGE_NAME = backend_image
CONTAINER_NAME = backend_container
FILE_WITH_ENV = ".env"
SHELL := /bin/bash
# GO = go
GO = go1.21.5

.PHONY:	run build lint format test cover 

#ARGS get in command string, example: make runbin ARGS="-m up -m-only"

up:
	docker compose up -d
	@echo wait for starting images...
# @sleep 10s
# @echo Create kafka topics:
# @docker-compose exec kafka kafka-topics \
# --create --topic messaggio --partitions 1 \
# --replication-factor 1 --bootstrap-server kafka:29092
# @docker-compose exec kafka kafka-topics \
# --create --topic messaggio-back --partitions 1 \
# --replication-factor 1 --bootstrap-server kafka:29092

docker-build: build
	docker build --tag mes1 .

down:
	docker compose down

run:
#@kafka-topics.sh --create --topic messaggio-back --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server localhost:9092
#@kafka-topics.sh --create --topic messaggio-back --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server localhost:9092
	@${GO} version
	@${GO} run ./cmd/${BINARY_NAME}/main.go

interrupt-main:
	pkill -2 main

interrupt-bin:
	pkill -2 ${BINARY_NAME}
	
# run: gen
# 	@${GO} version
# 	@${GO} run ./cmd/${BINARY_NAME}/main.go

# run binary file with arguments
runbin: gen
	@${GO} version
	@${GO} build -o ./build/binary/${BINARY_NAME} ./cmd/${BINARY_NAME}/main.go
	./build/binary/${BINARY_NAME} $(ARGS)

# for up DB by migration to top
db-up:
	@${GO} version
	@${GO} run ./cmd/${BINARY_NAME}/main.go -m-only up

# for reset DB by migration to 0
db-reset:
	@${GO} version
	@${GO} run ./cmd/${BINARY_NAME}/main.go -m-only reset

# for generation sqlc part
gen:
	@sqlc generate -f ./configs/sqlc/sqlc.yaml 

push:
	scp build/binary/messaggio_test anton@87.242.118.172:~/messaggio_test

build:
	@${GO} version
	${GO} build -o ./build/binary/${BINARY_NAME} ./cmd/${BINARY_NAME}/main.go

# build: gen
# 	@${GO} version
# 	${GO} build -o ./build/binary/${BINARY_NAME} ./cmd/${BINARY_NAME}/main.go

clean:
	@rm -R internal/sqlcGenerated ./build/binary/* logs/*

lint:
	@golangci-lint run ./... #-v
	@echo lint is OK

format:
	@gofmt -w */*/*.go
	@echo format is OK

test:
	go test ./... -v -cover 

cover:
	go test -covermode=count -coverpkg=./pkg/yandex -coverprofile ./tests/cover.out -v ./tests/pkg/yandex > /dev/null
	go tool cover -html ./tests/cover.out -o ./tests/cover.html > /dev/null
	@open ./tests/cover.html

docker-compose:
	docker compose up