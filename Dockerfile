# BUILD
FROM golang:1.22 AS builder

RUN go install github.com/a-h/templ/cmd/templ@v0.2.747

RUN dpkgArch="$(dpkg --print-architecture)"; \
    case "${dpkgArch##*-}" in \
        amd64) tailArch='linux-x64' ;; \
        arm64) tailArch='linux-arm64' ;; \
        *) echo >&2; echo >&2 "Uknown architecture ${dpkgArch}"; exit 1 ;; \
    esac; \
    wget https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.15/tailwindcss-${tailArch}; \
    chmod +x tailwindcss-${tailArch}; \
    mv tailwindcss-${tailArch} /usr/bin/tailwindcss;

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux make build

# DEPLOY
FROM alpine

WORKDIR /build

COPY --from=builder /app/main ./main

EXPOSE 9090

CMD ["./main"]
