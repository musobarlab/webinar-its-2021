# Builder
# ---------------------------------------------------
FROM golang:1.14-alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN go clean -modcache

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux \
        go build -ldflags '-w -s' -a -o go-codebase gitlab.com/Wuriyanto/go-codebase/cmd/go-codebase
# ---------------------------------------------------


# Final Image
# ---------------------------------------------------
FROM alpine

RUN apk update && \
    apk add --no-cache tzdata

ARG DIST_ENVIRONMENT=development

ENV TZ=Asia/Jakarta

WORKDIR /usr/src/app
ENV APP_PATH=/usr/src/app

COPY .env.dist.${DIST_ENVIRONMENT} ${APP_PATH}/.env
#COPY .env ${APP_PATH}/.env

COPY --from=builder /usr/src/app/go-codebase ${APP_PATH}/

EXPOSE 8000 8081

ENTRYPOINT ["/usr/src/app/go-codebase"]