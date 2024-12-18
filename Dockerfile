# Golangイメージをベースにする
FROM golang:1.21

# アプリケーションディレクトリの作成
WORKDIR /app

# 依存関係の追加
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# ポートの公開
EXPOSE 8080

# コンテナ起動時に実行するコマンド
CMD ["./main"]
