# アプリケーション名
APP_NAME := api-server

# ビルド出力先
BUILD_DIR := ./build

# メインパッケージ
MAIN_PACKAGE := ./cmd/api

# Swaggoの設定
SWAG_VERSION := v1.8.12
SWAG_INIT_OPTIONS := -g cmd/api/main.go --parseDependency --parseInternal

# デフォルトターゲット
.PHONY: all
all: clean build

# ドキュメント生成
.PHONY: docs
docs:
	@echo "Swaggoをインストールしています..."
	@go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)
	@echo "APIドキュメントを生成しています..."
	@$(shell go env GOPATH)/bin/swag init $(SWAG_INIT_OPTIONS)
	@echo "ドキュメント生成が完了しました。docs/ディレクトリを確認してください。"

# アプリケーションの実行
.PHONY: run
run:
	@echo "Swaggoをインストールしています..."
	@go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)
	@echo "APIドキュメントを生成しています..."
	@$(shell go env GOPATH)/bin/swag init $(SWAG_INIT_OPTIONS)
	@echo "ドキュメント生成が完了しました。docs/ディレクトリを確認してください。"
	@echo "アプリケーションを実行しています..."
	@go run $(MAIN_PACKAGE)

# アプリケーションのビルド
.PHONY: build
build:
	@echo "アプリケーションをビルドしています..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PACKAGE)
	@echo "ビルドが完了しました: $(BUILD_DIR)/$(APP_NAME)"

# テストの実行
.PHONY: test
test:
	@echo "テストを実行しています..."
	@go test -v ./...

# コードの静的解析
.PHONY: lint
lint:
	@echo "静的解析を実行しています..."
	@if [ -z "$$(which golangci-lint)" ]; then \
		echo "golangci-lintをインストールしています..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@golangci-lint run ./...

# 生成されたファイルの削除
.PHONY: clean
clean:
	@echo "生成されたファイルを削除しています..."
	@rm -rf $(BUILD_DIR)
	@echo "クリーンアップが完了しました。"

# Swagger UI の起動（オプション）
.PHONY: swagger-ui
swagger-ui:
	@echo "Swagger UIを起動しています..."
	@if [ -z "$$(which docker)" ]; then \
		echo "Dockerがインストールされていません。"; \
		exit 1; \
	fi
	@docker run --rm -p 8080:8080 -e SWAGGER_JSON=/docs/swagger.json -v $$(pwd)/docs:/docs swaggerapi/swagger-ui

# ヘルプ
.PHONY: help
help:
	@echo "利用可能なコマンド:"
	@echo "  make docs         - Swaggoを使用してAPIドキュメントを生成します"
	@echo "  make run          - アプリケーションを実行します"
	@echo "  make build        - アプリケーションをビルドします"
	@echo "  make test         - テストを実行します"
	@echo "  make lint         - コードの静的解析を実行します"
	@echo "  make clean        - 生成されたファイルを削除します"
	@echo "  make swagger-ui   - Swagger UIを起動します（Dockerが必要）"
	@echo "  make help         - このヘルプメッセージを表示します"
