# ================================================================================================
# api
# ================================================================================================
.PHONY: build run test clean

build: ## Build the API
	go build -o build/api/core/main cmd/api/core/main.go
	chmod +x build/api/core/main

run: build ## Run the API
	build/api/core/main -p 8080

test: ## Run tests
	go test ./...

clean:
	rm -rf build/api/core/main
	go clean

# ================================================================================================
# valkey
# ================================================================================================
.PHONY: valkey-up valkey-down valkey-restart valkey-rebuild valkey-logs valkey-cli

valkey-up: ## Start valkey
	docker compose up -d valkey

valkey-down: ## Stop valkey
	docker compose down valkey

valkey-restart: ## Restart valkey
	docker compose restart valkey

valkey-rebuild: ## Rebuild valkey
	docker compose down -v valkey && docker compose build --no-cache valkey && docker compose up -d valkey

valkey-logs: ## Show valkey logs
	docker compose logs -f valkey

valkey-cli: ## Run valkey CLI
	docker compose exec valkey valkey-cli

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
