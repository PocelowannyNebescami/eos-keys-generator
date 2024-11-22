FROM golang:1.22 AS fetch

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download && go mod verify;

FROM node:18-bookworm AS tailwind

RUN npm install -g tailwindcss@3.4.15;

WORKDIR /app

COPY . .

RUN tailwindcss -i ./cmd/web/assets/css/input.css -o ./output.css;

FROM ghcr.io/a-h/templ:latest AS templ

COPY --chown=65532:65532 . /app

WORKDIR /app

RUN ["templ", "generate"]

FROM fetch AS build

COPY --from=tailwind /app/output.css /app/cmd/web/assets/css/output.css

COPY --from=templ /app /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# DEPLOY
FROM alpine

WORKDIR /build

COPY --from=build /app/main ./main

EXPOSE 9090

ENTRYPOINT ["./main"]
