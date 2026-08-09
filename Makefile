.PHONY: api-dev web-dev oapi-check oapi-gen oapi-doc db-migrate-status db-migrate-up db-migrate-down

api-dev:
	go -C ./api run ./cmd/api

web-dev:
	pnpm --dir ./web dev

oapi-check:
	pnpm --package=@redocly/cli dlx redocly lint ./openapi/openapi.yml

oapi-gen:
	pnpm --package=@redocly/cli dlx redocly bundle ./openapi/openapi.yml -o ./openapi/openapi.bundle.yml
	go -C ./api tool oapi-codegen -config ./oapi-codegen.yaml ../openapi/openapi.bundle.yml

db-migrate-status:
	go -C ./api run ./cmd/migrate status
db-migrate-up:
	go -C ./api run ./cmd/migrate up
db-migrate-down:
	go -C ./api run ./cmd/migrate down
