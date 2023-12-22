BINARY_NAME=fhxparser
PORT=5101

.PHONY: run_front

build:
	@go build -o dist/${BINARY_NAME} ./cmd/api

api: build
	@env ./dist/${BINARY_NAME} -port=${PORT} &
	@echo "Backend running..."

stop:
	@-pkill -f ${BINARY_NAME}
	@echo "Backend stopped..."

clean: stop
	@go clean
	@rm -R ./dist

newrun: stop api

test:
	@go test ./...

install:
	@go mod tidy

front:
	cd cmd/frontend && npm run dev

run_app: api front

db:
	@echo "Starting database.."
	@docker compose up -d
	@echo "DB is started"

stop_db:
	@echo "Stoping database..."
	@docker compose down
	@echo "DB is stoped"


	