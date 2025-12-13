## Allow passing filenames as extra arguments, e.g.
##   make run-dev tests/step1/valid.json
## These extra words are captured into `ARGS` so make won't treat
## them as targets to build.
ARGS := $(filter-out build run run-dev,$(MAKECMDGOALS))

# Mark those extra args as phony so make doesn't try to build files
.PHONY: $(ARGS)
$(ARGS):
	@:

build:
	echo "Building the project..."
	go build -o build/json-parser

run: build
	./build/json-parser $(ARGS)

run-dev:
	go run . $(ARGS)