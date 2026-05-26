ARG GO_IMAGE=m.daocloud.io/docker.io/library/golang:1.25-alpine
ARG ALPINE_IMAGE=m.daocloud.io/docker.io/library/alpine:3.22

FROM ${GO_IMAGE} AS development

WORKDIR /app

ARG ALPINE_MIRROR=https://mirrors.aliyun.com/alpine
ARG GOPROXY=https://goproxy.cn,direct
ARG GOSUMDB=sum.golang.google.cn

ENV GOPROXY=${GOPROXY}
ENV GOSUMDB=${GOSUMDB}

RUN sed -i "s|https://dl-cdn.alpinelinux.org/alpine|${ALPINE_MIRROR}|g" /etc/apk/repositories
RUN apk add --no-cache bash ca-certificates curl git gzip mariadb-client openssh-client tar

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "./cmd/api"]

FROM development AS build

RUN CGO_ENABLED=0 go build -o /tmp/byml-api ./cmd/api
RUN CGO_ENABLED=0 go build -o /tmp/byml-restore ./cmd/restore

FROM ${ALPINE_IMAGE} AS production

WORKDIR /app

ARG ALPINE_MIRROR=https://mirrors.aliyun.com/alpine

RUN sed -i "s|https://dl-cdn.alpinelinux.org/alpine|${ALPINE_MIRROR}|g" /etc/apk/repositories
RUN apk add --no-cache bash ca-certificates curl git gzip mariadb-client openssh-client tar

COPY --from=build /tmp/byml-api /app/byml-api
COPY --from=build /tmp/byml-restore /app/byml-restore

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --retries=5 CMD curl -fsS http://localhost:8080/health >/dev/null || exit 1

CMD ["/app/byml-api"]
