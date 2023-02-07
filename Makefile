POSTGRES_DB_SCHEMA := public



postgres:
	docker run --name postgres12 --network q1-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root q1 

# 選擇DB: \connect simple123;  創建SCHEMA: CREATE SCHEMA v3
postgresTerminal:
	docker exec -it postgres12 psql -U root 

dropdb:
	docker exec -it postgres12 dropdb q1

migrateup:
	migrate -path migration -database "postgresql://root:admin@localhost:5432/q1?sslmode=disable&search_path=${POSTGRES_DB_SCHEMA}" -verbose up
migrateup1:
	migrate -path migration -database "postgresql://root:admin@localhost:5432/q1?sslmode=disable&search_path=${POSTGRES_DB_SCHEMA}" -verbose up 1
migratedown:
	migrate -path migration -database "postgresql://root:admin@localhost:5432/q1?sslmode=disable&search_path=${POSTGRES_DB_SCHEMA}" -verbose down
migratedown1:
	migrate -path migration -database "postgresql://root:admin@localhost:5432/q1?sslmode=disable&search_path=${POSTGRES_DB_SCHEMA}" -verbose down 1

test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mock -destination db/mock/store.go bank/db/sqlc Store
# 紀錄指令
# 創建升級 migration migrate create -ext sql -dir db/migration -seq add_users
# build image:  docker build -t q1:latest .
# build container: docker run --name q1 --network q1-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:admin@postgres12:5432/q1?sslmode=disable" q1:latest

.PHONY:postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test postgresTerminal server mock