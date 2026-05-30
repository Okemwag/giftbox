.PHONY: test migrate-status migrate-up migrate-down

test:
	go test ./...

migrate-status:
	go run ./cmd/migrate status

migrate-up:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down
