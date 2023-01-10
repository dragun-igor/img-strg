.PHONY: generate-proto
generate-proto:
	@cd proto && make generate

.PHONY: test
test:
	go test -v --cover internal/...

.PHONY: generate-mocks
generate-mocks:
	go generate ./...

.PHONY: deploy
deploy:
	sudo docker-compose up --build