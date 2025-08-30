#!/bin/bash
set -e

COMPOSE_FILE="docker-compose.nacos.yml"
WORK_DIR="$(dirname "$0")/../infrastructure/docker"

cd "$WORK_DIR"

case "$1" in
  start)
    echo "🚀 启动 Nacos..."
    docker compose -f $COMPOSE_FILE up -d
    echo "✅ Nacos 已启动"
    echo "👉 控制台: http://127.0.0.1:8848/nacos (账号 nacos / 密码 nacos)"
    ;;
  
  stop)
    echo "🛑 停止 Nacos..."
    docker compose -f $COMPOSE_FILE down
    echo "✅ Nacos 已停止"
    ;;
  
  restart)
    echo "🔄 重启 Nacos..."
    docker compose -f $COMPOSE_FILE down
    docker compose -f $COMPOSE_FILE up -d
    echo "✅ Nacos 已重启"
    echo "👉 控制台: http://127.0.0.1:8848/nacos (账号 nacos / 密码 nacos)"
    ;;
  
  clean)
    echo "🧹 停止并清理 Nacos（包括数据卷）..."
    docker compose -f $COMPOSE_FILE down -v
    echo "✅ Nacos 已彻底清理"
    ;;
  
  *)
    echo "用法: $0 {start|stop|restart|clean}"
    exit 1
    ;;
esac
