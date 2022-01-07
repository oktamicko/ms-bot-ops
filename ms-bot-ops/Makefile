NAME=ms-bot-ops
VERSION=0.0.2

build:
	@go build -o $(NAME)

build-prod:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags $(NAME)_$(VERSION) -ldflags '-w -extldflags "-static"' -o $(NAME) *.go

run-internal: build
	SERVICE_NAME=COBRANDING_SERVICE SERVICE_PORT=8087 ./$(NAME)

run-external: build
	SERVICE_NAME=COBRANDING_SERVICE SERVICE_PORT=8089 ./$(NAME)

run-profile: build
	@./$(NAME) -cpuprofile cpu.prof -memprofile mem.prof

clean:
	@rm -f $(NAME)
	@rm -rf vendor
	@rm -f cpu.prof
	@rm -f mem.prof


profile-mem:
	@go tool pprof --alloc_space ms-cobranding mem.prof

profile-cpu:
	@go tool pprof ms-cobranding cpu.prof

deps:
	git submodule update --init

deps-update:
	git submodule update --remote --merge

test:
	@go test -v *

lint:
	@go fmt bitbucket.org/bridce/ms-bot-ops/...

build-docker: build-prod
	docker build -t $(NAME):$(VERSION) .

run-docker:
	docker run --env-file .env -it $(NAME):$(VERSION)

.PHONY: deps lint build
