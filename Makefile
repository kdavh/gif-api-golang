NAME := giphy-api
PORT := 8080
HOST_PORT := 8080
APP_DIR := app

build:
	docker build -t ${NAME} .

run:
	docker run --name ${NAME}-container \
		--rm \
		-p ${HOST_PORT}:${PORT} \
		-it \
		-v `pwd`/${APP_DIR}:/go/src/GiphyApi \
		-w /go/src/GiphyApi \
		${NAME}
test:
	docker run --name ${NAME}-container-test \
		--rm \
		-it \
		-v `pwd`/${APP_DIR}:/go/src/GiphyApi \
		${NAME} go build && go test
sh:
	docker run --name ${NAME}-container-sh \
		--rm \
		-p ${HOST_PORT}:${PORT} \
		-it \
		-v `pwd`/${APP_DIR}:/go/src/GiphyApi \
		-w /go/src/GiphyApi \
		${NAME} /bin/bash
