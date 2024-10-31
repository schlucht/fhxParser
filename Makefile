BINARY_NAME=fhxparser
PORT=5101
DB_NAME=fhxdat.db

build:
	@go build -o dist/${BINARY_NAME} ./cmd/api

api: build
	@env ./dist/${BINARY_NAME} -port=${PORT} &
	@echo "Backend running..."

start: docker api

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

installdb:
	@brew install duckdb

rundb: stop
	@duckdb ./assets/database/${DB_NAME}

fronta:
	cd angular && npm start

docker:
	docker compose up -d





	