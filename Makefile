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
