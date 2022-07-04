all: build run
build:
	docker build --tag=test-scripts:latest -f Dockerfile.test . --build-arg test_dir=scripts
	docker build --tag=test-maths:latest -f Dockerfile.test . --build-arg test_dir=maths
	docker build --tag=go-scripts:latest -f Dockerfile .
run:
	echo "\n\n ###   TEST-SCRIPTS   #### \n\n"
	docker run --name tests -it --rm test-scripts:latest
	echo "\n\n ###   TEST-MATHS   #### \n\n"
	docker run --name tests -it --rm test-maths:latest
	docker run --name scripts -it --rm go-scripts:latest
