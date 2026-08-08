.PHONY: api-dev web-dev

api-dev:
	cd ./api &&\
	go run ./cmd/api

web-dev:
	cd ./web &&\
	pnpm dev
