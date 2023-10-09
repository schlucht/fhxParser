BINARY_NAME=fhxparser
PORT=5101

.PHONY: run_front

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

newrun: stop run

test:
	@go test ./...

install:
	@go mod tidy

run_front:
	cd cmd/frontend && npm run dev

run_app: run run_front