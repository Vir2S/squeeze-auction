dc-run:
	docker-compose up -d

swag-init:
	swag init -g cmd/main.go

run:
	go run cmd/main.go