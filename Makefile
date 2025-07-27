# ================================================================================================
# api
# ================================================================================================
# docker
.PHONY: up down logs

up: ## Start containers
	docker compose up -d core-api redis

down: ## Stop containers
	docker compose down -v core-api redis

logs: ## Show logs
	docker compose logs -f core-api

# ================================================================================================
# batch
# ================================================================================================
.PHONY: batch-set batch-sadd batch-rpush

batch-set: ## Run batch SET (use FLAGS=-large=true for large data, default: small)
	REDIS_HOST=localhost go run cmd/batch/set/*.go $(FLAGS)

batch-sadd: ## Run batch SADD (use FLAGS=-large=true for large data, default: small)
	REDIS_HOST=localhost go run cmd/batch/sadd/*.go $(FLAGS)

batch-rpush: ## Run batch RPUSH (use FLAGS=-large=true for large data, default: small)
	REDIS_HOST=localhost go run cmd/batch/rpush/*.go $(FLAGS)

# ================================================================================================
# redis
# ================================================================================================
.PHONY: redis-cli redis-up redis-down

redis-up: ## Start redis container
	docker compose up -d redis --build

redis-down: ## Stop redis container
	docker compose down -v redis

redis: ## Start session with redis-cli
	docker compose exec redis redis-cli

redis-cli: ## Run redis command with arguments (usage: make redis-cli CMD="DBSIZE")
	docker compose exec redis redis-cli $(CMD)

redis-flush: ## Flush all data from Redis
	docker compose exec redis redis-cli FLUSHALL

# ================================================================================================
# aws (elasticache)
# ================================================================================================
.PHONY: elasticache-valkey-versions elasticache-valkey-versions-only valkey-docker-tags

elasticache-valkey-versions: ## Show valkey versions
	aws elasticache describe-cache-engine-versions \
		--profile $(AWS_PROFILE) \
		--engine "Valkey" | jq .

elasticache-valkey-versions-only: ## Show valkey versions only
	aws elasticache describe-cache-engine-versions \
		--profile $(AWS_PROFILE) \
		--engine "Valkey" \
		--query "CacheEngineVersions[].EngineVersion" | jq .

valkey-docker-tags: ## Show valkey docker tags
	curl -s "https://hub.docker.com/v2/repositories/valkey/valkey/tags?page_size=100" | jq -r '.results[].name' | grep -E '^8\.1'

# =================================================================
# general
# =================================================================
.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
