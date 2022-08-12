FROM golang

WORKDIR /app

COPY . .

ENV GOPATH=/

EXPOSE 8080

RUN go mod download
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go build -o notes cmd/main.go

CMD ["./notes"]