include .env

.PHONY: up down mysql swag lint lint-fix test test-cover cover-avg

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

lint-fix:
	docker exec -it frescos-api sh -c "golangci-lint run ./... --fix"

test:
	go run gotest.tools/gotestsum@latest --format dots $(or $(filter-out $@,$(MAKECMDGOALS)),./...)

test-cover:
	go test -cover -coverprofile=./tmp/coverage.out ./internal/controllers/... ./internal/service/... \
	&& sed -i '' '/_mock.go/d' ./tmp/coverage.out

cover-avg:
	go tool cover -func=./tmp/coverage.out | grep total | awk '{print $3}'
