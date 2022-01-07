protoc:
	docker build -t protocon .
	docker run --rm -v $(shell pwd)/pb:/tmp protocon /bin/sh -c "cp /home/*.pb.go /tmp/"

clean:
	docker rmi protocon

.PHONY: build clean