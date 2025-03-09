.DEFAULT_GOAL := help

# Dockerイメージのタグ
DOCKER_TAG := latest

# 本番環境用のDockerイメージをビルド
build:
	docker build -t budougumi0617/gotodo:${DOCKER_TAG} \
		--target deploy ./

# ローカル開発用のDockerイメージをビルド
build-local:
	docker compose build --no-cache

# Docker Composeで開発環境を起動
up:
	docker compose up

# Docker Composeのコンテナを停止
down:
	docker compose down

# ログをリアルタイムで表示
logs:
	docker compose logs -f

# 実行中のコンテナのステータスを確認
ps:
	docker compose ps

help:
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
