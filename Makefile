# ================================================================================================
# api
# ================================================================================================
.PHONY: build run test clean

build:
	go build -o build/api/core/main cmd/api/core/main.go
	chmod +x build/api/core/main

run: build
	build/api/core/main -p 8080

test:
	go test ./...

clean:
	rm -rf build/api/core/main
	go clean

# ================================================================================================
# valkey
# ================================================================================================
.PHONY: valkey-up valkey-down valkey-restart valkey-rebuild valkey-logs valkey-cli

valkey-up:
	docker compose up -d valkey

valkey-down:
	docker compose down valkey

valkey-restart:
	docker compose restart valkey

valkey-rebuild:
	docker compose down -v valkey && docker compose build --no-cache valkey && docker compose up -d valkey

valkey-logs:
	docker compose logs -f valkey

valkey-cli:
	docker compose exec valkey valkey-cli

# ================================================================================================
# aws (elasticache)
# ================================================================================================
.PHONY: elasticache-valkey-versions elasticache-valkey-versions-only valkey-docker-tags

elasticache-valkey-versions:
	aws elasticache describe-cache-engine-versions \
		--profile $(AWS_PROFILE) \
		--engine "Valkey" | jq .

elasticache-valkey-versions-only:
	aws elasticache describe-cache-engine-versions \
		--profile $(AWS_PROFILE) \
		--engine "Valkey" \
		--query "CacheEngineVersions[].EngineVersion" | jq .

valkey-docker-tags:
	curl -s "https://hub.docker.com/v2/repositories/valkey/valkey/tags?page_size=100" | jq -r '.results[].name' | grep -E '^8\.1'
