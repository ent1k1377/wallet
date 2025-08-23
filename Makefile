.PHONY: mig-up mig-down

include configs/.env

DSN = "postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${DB_HOST}:${DB_PORT}/${POSTGRES_DB}?sslmode=disable"

mig-up:
	goose -dir migrations/ postgres ${DSN} up

mig-down:
	goose -dir migrations/ postgres ${DSN} down

up:
	docker compose --env-file configs/.env -f deployments/compose.yaml up

build:
	docker compose --env-file configs/.env -f deployments/compose.yaml build