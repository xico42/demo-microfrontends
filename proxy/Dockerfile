FROM node:20-alpine AS layout-builder
WORKDIR /build
COPY layout/package.json ./
RUN npm install
COPY layout/ ./
RUN npm run build

FROM golang:1.25-alpine AS caddy-builder
WORKDIR /build
COPY proxy/src/go.mod proxy/src/go.sum ./
RUN go mod download
COPY proxy/src/ ./
RUN CGO_ENABLED=0 go build -o /caddy ./main.go

FROM alpine:3.21
COPY --from=caddy-builder /caddy /usr/bin/caddy
COPY --from=layout-builder /build/dist/ /srv/layout/
COPY proxy/Caddyfile /etc/caddy/Caddyfile
CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]
