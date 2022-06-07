all: build

VERSION := 0.0.1

# Primary targets
build: main

test:
	go test

deploy: deploy.yaml
	kubectl apply -f deploy.yaml

dist: main
	tar cvf dist/server-$(VERSION).tar.gz main
	make docker
	
clean:
	rm -rf dist/*

# Rest of the stuff

dev: main
	./main

docker: main Dockerfile
	docker build .

main: main.go
	go build main.go
