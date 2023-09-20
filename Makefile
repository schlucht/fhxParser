BINARY_NAME=fhxparser
PORT=5101

build:
	@go build -o dist/${BINARY_NAME} ./cmd/api

run: build
	@env ./dist/${BINARY_NAME} -port=${PORT} &
	@echo "Backend running..."

stop:
	@-pkill -f ${BINARY_NAME}
	@echo "Backend stopped..."

clean: stop
	@go clean
	@rm -R ./dist

test:
	@go test ./...

install:
	@go mod tidy