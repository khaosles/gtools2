.PHONY: build test run proto clean docker-build

NAME=app
DB_FILE=app.db.bin
PB_DIR=./pb

# version settings
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
PROJECT?=github.com/khaosles/gtools2

build:
	go build -ldflags "-X ${PROJECT}/pkg/handler.Release=${RELEASE} \
	-X ${PROJECT}/pkg/handler.Commit=${COMMIT} -X ${PROJECT}/pkg/handler.BuildTime=${BUILD_TIME}" \
	-o ${NAME} "${PROJECT}/cmd/server"

test:
	go test -race github.com/gerlacdt/db-key-value-store/...

run: build
	PORT=8080 FILENAME=${DB_FILE} ./app

proto:
	protoc -I ${PB_DIR} ${PB_DIR}/db.proto --go_out=${PB_DIR}

clean:
	rm -f ${NAME} ${DB_FILE} ./pkg/db/db.test.bin

docker-build:
	GOOS=linux go build -o ${NAME} "${PROJECT}/cmd/server"
	docker build -t gerlacdt/db-key-value-store:latest .
	rm -f ${NAME}