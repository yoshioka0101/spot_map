# デプロイ用のビルドコンテナ
FROM golang:1.23.3-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -trimpath -ldflags "-w -s" -o app

# デプロイ用のランタイムコンテナ
FROM debian:bullseye-slim AS deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ローカル開発環境用のホットリロード環境
FROM golang:1.23.3 AS dev

WORKDIR /app

# airのインストール
RUN go install github.com/air-verse/air@latest

CMD ["air"]
