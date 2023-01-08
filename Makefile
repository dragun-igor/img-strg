.PHONY: generate-proto
generate-proto:
	@cd proto && make generate

.PHONY: test
test:
	cd test && go test -v

.PHONY: deploy
deploy:
	sudo docker-compose up --build