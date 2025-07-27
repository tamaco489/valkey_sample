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
