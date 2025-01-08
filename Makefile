include .env

.PHONY: up down mysql

up:
	docker-compose up -d

down:
	docker-compose down

mysql:
	docker exec -it frescos-db mysql -u${DB_USER} -p${DB_PASS} ${DB_NAME}
