all: build

# Primary targets
build: main

test:
	go test

dist: main
	tar cvf dist/release.tar.gz main
	
clean:
	rm -rf dist/*

# Rest of the stuff

dev: main
	./main

docker: main Dockerfile
	docker build .

main: main.go
	go build main.go