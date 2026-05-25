FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache bash ca-certificates git gzip mariadb-client openssh-client tar

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "./cmd/api"]
