BINARY_NAME=golarisApp

build:
	@go mod vendor
	@echo "Building golaris..."
	@go build -o tmp/$(BINARY_NAME) .
	@echo "Golaris built!"

run: build
	@echo "Running golaris app..."
	@./tmp/$(BINARY_NAME) &
	@echo "Golaris app running!"

clean:
	@echo "Cleaning up..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned up!"

test:
	@echo "Testing golaris..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping golaris app..."
	@-pkill -SIGTERM -f "./tmp/$(BINARY_NAME)"
	@echo "Golaris app stopped!"

restart: stop start
