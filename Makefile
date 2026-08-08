.PHONY: api-dev web-dev openapi-check openapi-gen

api-dev:
	go -C ./api run ./cmd/api

web-dev:
	pnpm --dir ./web dev

openapi-check:
	pnpm --package=@redocly/cli dlx redocly lint ./openapi/openapi.yml

openapi-gen:
	pnpm --package=@redocly/cli dlx redocly bundle ./openapi/openapi.yml -o ./openapi/openapi.bundle.yml
	go -C ./api tool oapi-codegen -config ./oapi-codegen.yaml ../openapi/openapi.bundle.yml