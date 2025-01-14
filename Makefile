include .env

.PHONY: up down mysql swag lint

up:
	docker-compose up -d

down:
	docker-compose down

mysql:
	docker exec -it frescos-db mysql -u${DB_USER} -p${DB_PASS} ${DB_NAME}

swag:
	docker exec -it frescos-api sh -c "swag init -d cmd --parseDependency --parseInternal --parseDepth 4 -o swagger/docs"

lint:
	docker exec -it frescos-api sh -c "golangci-lint run ./..."