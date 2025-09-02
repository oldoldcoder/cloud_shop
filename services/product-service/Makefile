.PHONY: build run test clean deps swagger

# 构建应用
build:
	go build -o bin/product-service cmd/product-service/main.go

# 运行应用
run:
	go run cmd/product-service/main.go

# 运行测试
test:
	go test ./...

# 清理构建文件
clean:
	rm -rf bin/

# 安装依赖
deps:
	go mod tidy
	go mod download

# 生成 Swagger 文档
swagger:
	swag init -g cmd/product-service/main.go

# 构建 Docker 镜像
docker-build:
	docker build -t product-service .

# 运行 Docker 容器
docker-run:
	docker run -p 9082:9082 product-service

# 开发模式（包含热重载）
dev:
	air

# 安装 air（热重载工具）
install-air:
	go install github.com/cosmtrek/air@latest
