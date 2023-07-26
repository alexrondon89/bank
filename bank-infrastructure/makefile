migrateup:
	migrate -path ./database -database "postgres://postgres:mysecretpassword@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./database -database "postgres://postgres:mysecretpassword@localhost:5432/bank?sslmode=disable" -verbose down

compose:
	docker-compose up -d