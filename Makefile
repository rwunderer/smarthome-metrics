BASE_REPO 	:= rwunderer
APP 		:= smarthome-metrics
IMAGE 		:= $(BASE_REPO)/$(APP)

.PHONY: docker-build
docker-build:
	docker build -t $(IMAGE) . -f build/package/Dockerfile

.PHONY: docker-run
docker-run: docker-build
	docker run --rm -it $(IMAGE)

.PHONY: docker-push
docker-push: docker-build
	docker push $(IMAGE)

.PHONY: test
test:
	go test ./internal/pkg/config/

.PHONY: build
build:
	go build ./cmd/$(APP)

.PHONY: fmt
fmt:
	@find . -name "*.go" -exec go fmt {} \;

.PHONY: run
run:
	go run ./cmd/$(APP)/main.go

.PHONY: clean
clean:
	rm -f $(APP)

# vim: noexpandtab
