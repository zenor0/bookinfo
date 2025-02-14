#!/bin/bash

# 检查环境变量
if [ -z "$POSTGRES_HOST" ]; then
    POSTGRES_HOST="localhost"
fi

if [ -z "$POSTGRES_PORT" ]; then
    POSTGRES_PORT="5432"
fi

if [ -z "$POSTGRES_USER" ]; then
    POSTGRES_USER="postgres"
fi

if [ -z "$POSTGRES_PASSWORD" ]; then
    POSTGRES_PASSWORD="postgres"
fi

# 等待 PostgreSQL 启动
echo "Waiting for PostgreSQL to start..."
until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -c '\q'; do
    echo "PostgreSQL is unavailable - sleeping"
    sleep 1
done

echo "PostgreSQL is up - executing init script"

# 执行初始化脚本
PGPASSWORD=$POSTGRES_PASSWORD psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -f init.sql

echo "Database initialization completed" 