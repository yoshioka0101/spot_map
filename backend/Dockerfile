# デプロイ用のビルドコンテナ
FROM golang:1.23.3-bullseye AS deploy-builder

WORKDIR /app

# モジュールの依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . ./
RUN go build -trimpath -ldflags "-w -s" -o app ./cmd/api

# デプロイ用のランタイムコンテナ
FROM debian:bullseye-slim AS deploy

WORKDIR /app

RUN apt-get update

COPY --from=deploy-builder /app/app /app/app

CMD ["/app/app"]

# ローカル開発環境用のホットリロード環境
FROM golang:1.23.3 AS dev

WORKDIR /app

# airのインストール
RUN go install github.com/air-verse/air@latest

# air を実行
CMD ["air"]
