build:
	docker build -t notes .

run:
	docker run --name notes -p 8080:8080 -it notes

remove:
	docker rm -f notes
	docker rmi notes

migrate:
	migrate -path ./migrations/postgres/schemas -database postgres://postgres:secret@localhost:5436/notes?sslmode=disable up

swagger:
	swag init -g cmd/main.go

postgres:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=notes -d -v notes:/var/lib/postgresql/data -p 5436:5432 postgres