FROM golang:1.22-alpine AS builder

RUN apk add build-base
RUN apk update && apk add git

WORKDIR /ai-powered-chatbot-project-monorepo

COPY ./go.mod ./
COPY ./go.sum ./
COPY ./search-engine-service ./search-engine-service
COPY ./common ./common
COPY ./gokit ./gokit

RUN go mod download

RUN cd /ai-powered-chatbot-project-monorepo/search-engine-service && go build -o search-engine-service ./cmd/main.go

EXPOSE 8686 9090

FROM alpine:latest
RUN apk add ca-certificates libstdc++
COPY --from=builder /ai-powered-chatbot-project-monorepo/search-engine-service /usr/local/bin/
ENTRYPOINT ["search-engine-service"]





