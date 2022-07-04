all: build run
build:
	docker build --tag=go-scripts-tests:latest -f Dockerfile.test .
	docker build --tag=go-scripts:latest -f Dockerfile .
run:
	docker run --name tests -it --rm go-scripts-tests:latest
	docker run --name scripts -it --rm go-scripts:latest
