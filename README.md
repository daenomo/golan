# golan

軽量な Go アプリケーションのサンプルリポジトリ（`devcontainer` を使った Docker 開発環境 含む）。

## 概要

このリポジトリは、Docker と `devcontainer` を使ってローカルで開発できる Go アプリケーションの最小構成を提供します。

## 特長

- Docker と `docker-compose` で開発環境を再現可能
- シンプルな `main.go` を使ったサンプルアプリ
- VS Code の `devcontainer` による一貫した開発体験

## 前提条件

- Docker
- docker-compose
- （ローカルで直接ビルドする場合）Go 1.20+ がインストールされていること

## 使い方（クイックスタート）

コンテナで実行する（推奨）:

```bash
docker compose up --build
```

アプリケーションをローカルでビルドして実行する:

```bash
go build -o bin/golan ./app
./bin/golan
```

VS Code で `devcontainer` を使う場合: 「Reopen in Container」を選んでください。

## 開発

- 変更を加えたらコンテナを再ビルドするか、ローカルで再ビルドしてください。
- 単体テストがあれば `go test ./...` を使います。

## ディレクトリ構成

```
.
├─ app/            # アプリケーションのソース (main.go)
├─ Dockerfile
├─ docker-compose.yml
├─ .devcontainer/  # VS Code devcontainer 設定
└─ README.md
```

## 貢献

Issue や Pull Request を歓迎します。変更点は簡潔に説明してください。

## ライセンス

特に指定がない場合はリポジトリ所有者にお問い合わせください。
