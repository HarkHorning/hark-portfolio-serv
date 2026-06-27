.PHONY: up down restart logs build clean help

help:
	@echo "Home Server Management:"
	@echo "  make up      - Start the site in the background"
	@echo "  make down    - Stop the site"
	@echo "  make logs    - View logs for all services"
	@echo "  make build   - Rebuild Docker images"
	@echo "  make clean   - Stop and remove all data (CAREFUL: wipes DB)"

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

logs:
	docker compose logs -f

build:
	docker compose build --no-cache

clean:
	docker compose down -v