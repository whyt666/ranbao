# 定义变量
DOCKER_COMPOSE_FILE := docker-compose.yml
DB_SERVICE := db
DB_ROOT_PASSWORD := 1520226681wW
DB_NAME := ranbao
DEPLOY_DIR := d:/golang_project/ranbao/deploy
SQL_FILES := $(wildcard $(DEPLOY_DIR)/*.sql)

# 启动 Docker 服务
start:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# 停止 Docker 服务
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# 执行 SQL 文件
migrate: start
	@for file in $(SQL_FILES); do \
		docker exec -i $(shell docker-compose -f $(DOCKER_COMPOSE_FILE) ps -q $(DB_SERVICE)) mysql -u root -p$(DB_ROOT_PASSWORD) $(DB_NAME) < $$file; \
	done

# 默认目标
all: migrate

.PHONY: start stop migrate all