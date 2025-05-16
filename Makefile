build: build-windows-amd64 build-linux-amd64 build-linux-arm64
build-%:
	$(eval export CGO_ENABLED = 0)
	$(eval OSARCH = $(subst -, ,$*))
	$(eval export GOOS = $(word 1,$(OSARCH)))
	$(eval export GOARCH = $(word 2,$(OSARCH)))
	@echo Building $*
	go build -ldflags "-s -w" -trimpath -o out/$(GOOS)-$(GOARCH)/
	cp LICENSE COPYRIGHT NOTICE README.md out/$(GOOS)-$(GOARCH)/
