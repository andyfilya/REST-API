build:
	docker-compose build restapi
run:
	docker-compose up restapi
migrate:
	migrate -path ./schema -database 'postgres://postgres:root@localhost:5432/users_inf?sslmode=disable' up
swag:
	swag init -g cmd/REST-API/main.go