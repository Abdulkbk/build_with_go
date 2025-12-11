build:
	echo "Building the project..."
	go build -o build/json-parser

run: build
	./build/json-parser

run-dev:
	go run . $(ARGS)