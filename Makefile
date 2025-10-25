# Makefile for the 'ai' project

# Default target
all: build

# Build the application
build:
	go build -o ai .

# Run the application with arguments
# Usage: make run ARGS="your arguments here"
run:
	go run . "$(ARGS)"

# Clean the project
clean:
	rm -f ai
