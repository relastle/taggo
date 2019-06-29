.PHONY: main
main:
	go build

.PHONY: docker
docker:
	docker build -t relastle/taggo:0.1.0 -f docker/Dockerfile .
