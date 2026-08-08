.PHONY: api-dev web-dev openapi-check

api-dev:
	cd ./api &&\
	go run ./cmd/api

web-dev:
	cd ./web &&\
	pnpm dev

openapi-check:
	pnpm --package=@redocly/cli dlx redocly lint ./openapi/openapi.yml
