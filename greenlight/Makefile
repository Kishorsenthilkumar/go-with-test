.PHONY:	run/api
run/api:
	@go run ./cmd/api -db-dsn=${GREENLIGHT_DB_DSN}

.PHONY: db/psql
db/psql:
	psql ${GREENLIGHT_DB_DSN}

.PHONY:	db/migration/new
db/migration/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext .sql -dir ./migrations ${name}

.PHONY:	db/migration/up
db/migration/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up

.PHONY:	audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

.PHONY:	build/api
build/api:
	@echo 'Building cmd/api...'
	go build -o=./bin/api ./cmd/api