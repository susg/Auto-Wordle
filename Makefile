IMAGE_NAME=autowordle
.PHONY: build run test

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -it $(IMAGE_NAME)

test:
	docker run -it $(IMAGE_NAME) go test -v ./...
