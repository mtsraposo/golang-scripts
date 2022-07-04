all: build run
build:
	docker build --tag=go-scripts:latest -f Dockerfile .
	docker build --tag=go-scripts-tests:latest -f Dockerfile.test .
run:
	docker run --name scripts -it go-scripts:latest
	docker run --name tests -it go-scripts-tests:latest