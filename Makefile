build:
	docker build -t notes .
run:
	docker run --name notes -p 8080:8080 -it notes
remove:
	docker rm -f notes
	docker rmi notes

migrate:
	migrate -path ./migrations/postgres/schemas -database postgres://postgres:secret@localhost:5000/notes?sslmode=disable up

swagger:
	swag init -g cmd/main.go

postgres:
	docker run --name postgres -p 5000:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres
	docker exec postgres /bin/sh -c "echo Welcome to postgres container, run: psql -h localhost -U postgres, if you want connect to the psql shell"
	docker exec -it postgres /bin/sh
db:
	docker exec postgres createdb --username=postgres --owner=postgres notes