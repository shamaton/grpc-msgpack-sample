protoc:
	git submodule update --init
	docker build -t protocon .
	docker run --rm -v $(shell pwd)/pb:/tmp protocon /bin/sh -c "cp /home/*.pb.go /tmp/"

client:
	go run ./cmd/client/

server:
	go run ./cmd/server/

clean:
	docker rmi protocon

.PHONY: protoc client server clean