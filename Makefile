
migrateup:
	migrate -path db/migration -database "postgresql://root:pass123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:pass123@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup, migratedown, sqlc
