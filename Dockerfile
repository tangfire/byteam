FROM node:22-alpine AS frontend-build

WORKDIR /app

ARG ALPINE_MIRROR=https://mirrors.aliyun.com/alpine
ARG NPM_REGISTRY=https://registry.npmmirror.com

RUN sed -i "s|https://dl-cdn.alpinelinux.org/alpine|${ALPINE_MIRROR}|g" /etc/apk/repositories
RUN npm config set registry "${NPM_REGISTRY}"

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM nginx:1.27-alpine

ARG ALPINE_MIRROR=https://mirrors.aliyun.com/alpine

RUN sed -i "s|https://dl-cdn.alpinelinux.org/alpine|${ALPINE_MIRROR}|g" /etc/apk/repositories
RUN apk add --no-cache curl

COPY deploy/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=frontend-build /app/dist /usr/share/nginx/html

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=5s --retries=5 CMD curl -fsS http://localhost/ >/dev/null || exit 1
