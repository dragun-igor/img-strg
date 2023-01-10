.PHONY: generate-proto
generate-proto:
	@cd proto && make generate

.PHONY: test
test:
	cd test && go test -v

.PHONY: generate-mocks
generate-mocks:
	rm -rf internal/server/service/mocks && go generate internal/server/service/interface.go
	rm -rf internal/pkg/storage/mocks && go generate internal/pkg/storage/interface.go

.PHONY: deploy
deploy:
	sudo docker-compose up --build