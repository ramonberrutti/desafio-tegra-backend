FROM golang AS build-env

RUN go get github.com/rs/cors

WORKDIR /go/src/github.com/ramonberrutti/desafio-tegra-backend
COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o desafio-tegra-backend

FROM alpine
WORKDIR /app

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true && rm -rf /var/cache/apk/*

COPY --from=build-env /go/src/github.com/ramonberrutti/desafio-tegra-backend/desafio-tegra-backend /app/

EXPOSE 8080
ENTRYPOINT ./desafio-tegra-backend