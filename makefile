
# Variables
APP_NAME=Blog-CMS
BUILD_DIR=bin
SRC_DIR=.

# Build executable
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# Run the application
run: build
	$(BUILD_DIR)/$(APP_NAME)


# goose local
DB_HOST=127.0.0.1
DB_PORT=3308
DB_USER=blogcms
DB_PASSWORD=S@ng0905257554
DB_NAME=blogcms
DB_DRIVER=mysql
DSN=$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?charset=utf8mb4&parseTime=True&loc=Local

# Thư mục migrations
GOOSE_DBSTRING ?= "root:root1234@tcp(127.0.0.1:33306)/shopdevgo"
MIGRATIONS_DIR=migrations
GOOSE_DRIVER ?= mysql

# Cài đặt Goose nếu chưa có
install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

# Lệnh chạy migrations
migrate-up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DSN)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DSN)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DSN)" status

migrate-create:
	@if [ -z "$(NAME)" ]; then echo "Missing NAME. Use 'make migrate-create NAME=your_migration_name'"; exit 1; fi
	goose -dir $(MIGRATIONS_DIR) create $(NAME) sql

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs