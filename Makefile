.env:
	cp .env.example .env

.PHONY: deps
deps: .env
	docker-compose run --rm golang go mod tidy

.PHONY: test
test: .env
	docker-compose run --rm golang go test -v ./...

.PHONY: testIncludeInt
testIncludeInt: .env
	$(MAKE) startDB
	docker-compose run --rm golang go test -tags=integration -v ./...
	$(MAKE) stop

.PHONY: build
build: .env
	docker-compose build app
	docker image prune -f --filter label=stage=build

.PHONY: startDB
startDB: .env
	docker-compose up -d postgres

.PHONY: quickStart
quickStart: .env
	$(MAKE) startDB
	docker-compose up --build app

.PHONY: stop
stop: .env
	docker-compose down

.PHONY: fmt
fmt: .env
	docker-compose run --rm golang go fmt ./...

.PHONY: genMocks
genMocks: .env
	rm -rf mocks/*
	docker-compose run --rm mockery --all
