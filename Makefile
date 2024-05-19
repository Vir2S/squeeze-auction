dc-run:
	docker-compose up -d

dc-stop:
	docker-compose down

swag-init:
	swag init -g cmd/main.go

run:
	go run cmd/main.go
