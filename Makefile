BINARY_NAME=fhxparser
PORT=5101
DB_NAME=fhxdat.db

build:
	@go build -o dist/${BINARY_NAME} ./cmd/api

api: build
	@env ./dist/${BINARY_NAME} -port=${PORT} &
	@echo "Backend running..."

stop:
	@-pkill -f ${BINARY_NAME}
	@echo "Backend stopped..."

restart: stop api
	clear

clean: stop
	@go clean
	@rm -R ./dist

newrun: stop api

test:
	@go test ./...

install:
	@go mod tidy

run: api

installdb:
	@brew install duckdb

rundb: stop
	@duckdb ./assets/database/${DB_NAME}





	