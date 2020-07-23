#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------
OUTDIR ?= _output
PROJECT ?= service-mesh-hub

DOCKER_REPO ?= soloio
PROJECT_IMAGE ?= $(DOCKER_REPO)/service-mesh-hub

SOURCES := $(shell find . -name "*.go" | grep -v test.go)
RELEASE := "true"
ifeq ($(TAGGED_VERSION),)
	TAGGED_VERSION := $(shell git describe --tags --dirty --always)
	RELEASE := "false"
endif

VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)
.PHONY: print-version
print-version:
ifeq ($(TAGGED_VERSION),)
	exit 1
endif
	echo $(VERSION)

LDFLAGS := "-X github.com/solo-io/$(PROJECT)/pkg/common/version.Version=$(VERSION)"
GCFLAGS := all="-N -l"

print-info:
	@echo RELEASE: $(RELEASE)
	@echo TAGGED_VERSION: $(TAGGED_VERSION)
	@echo VERSION: $(VERSION)

#----------------------------------------------------------------------------------
# Code Generation
#----------------------------------------------------------------------------------
.PHONY: fmt
fmt:
	goimports -w $(shell ls -d */ | grep -v vendor)

.PHONY: mod-download
mod-download:
	go mod download

# Dependencies for code generation
.PHONY: install-go-tools
install-go-tools: mod-download
	go install istio.io/tools/cmd/protoc-gen-jsonshim
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/golang/protobuf/protoc-gen-go
	go install github.com/solo-io/protoc-gen-ext
	go install github.com/golang/mock/mockgen
	go install golang.org/x/tools/cmd/goimports
	go install github.com/onsi/ginkgo/ginkgo

# Call all generated code targets
.PHONY: generated-code
generated-code: operator-gen \
				manifest-gen \
				go-generate \
				generated-reference-docs \
				fmt
	go mod tidy

#----------------------------------------------------------------------------------
# Go generate
#----------------------------------------------------------------------------------

# Run go-generate on all sub-packages
go-generate:
	go generate -v ./...

#----------------------------------------------------------------------------------
# Operator Code Generation
#----------------------------------------------------------------------------------

# Generate Operator Code
.PHONY: operator-gen
operator-gen:
	go run -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) codegen/generate.go

#----------------------------------------------------------------------------------
# Docs Code Generation
#----------------------------------------------------------------------------------

# TODO(ilackarms): broken, fix
# Generate Reference documentation
.PHONY: generated-reference-docs
generated-reference-docs:
#	go run docs/generate_reference_docs.go

#----------------------------------------------------------------------------------
# Build
#----------------------------------------------------------------------------------

.PHONY: build-all-images
build-all-images: service-mesh-hub-image

#----------------------------------------------------------------------------------
# Build service-mesh-hub controller + cli images
#----------------------------------------------------------------------------------

# for local development only; to build docker image, use service-mesh-hub-linux-amd-64
.PHONY: service-mesh-hub
service-mesh-hub: $(OUTDIR)/service-mesh-hub
$(OUTDIR)/service-mesh-hub: $(SOURCES)
	go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/service-mesh-hub/main.go

.PHONY: service-mesh-hub-linux-amd64
service-mesh-hub-linux-amd64: $(OUTDIR)/service-mesh-hub-linux-amd64
$(OUTDIR)/service-mesh-hub-linux-amd64: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/service-mesh-hub/main.go

.PHONY: meshctl-linux-amd64
meshctl-linux-amd64: $(OUTDIR)/meshctl-linux-amd64
$(OUTDIR)/meshctl-linux-amd64: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/meshctl/main.go

.PHONY: meshctl-darwin-amd64
meshctl-darwin-amd64: $(OUTDIR)/meshctl-darwin-amd64
$(OUTDIR)/meshctl-darwin-amd64: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/meshctl/main.go

.PHONY: meshctl-windows-amd64
meshctl-windows-amd64: $(OUTDIR)/meshctl-windows-amd64.exe
$(OUTDIR)/meshctl-windows-amd64.exe: $(SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ cmd/meshctl/main.go

.PHONY: build-cli
build-cli: meshctl-linux-amd64 meshctl-darwin-amd64 meshctl-windows-amd64

.PHONY: install-cli
install-cli:
	go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o ${GOPATH}/bin/meshctl cmd/meshctl/main.go

# build image with service-mesh-hub binary
# this is an alternative to using operator-gen to build the image
.PHONY: service-mesh-hub-image
service-mesh-hub-image: service-mesh-hub-linux-amd64
	cp $(OUTDIR)/service-mesh-hub-linux-amd64 build/service-mesh-hub/ && \
	docker build -t $(PROJECT_IMAGE):$(VERSION) build/service-mesh-hub/
	rm build/service-mesh-hub/service-mesh-hub-linux-amd64

#----------------------------------------------------------------------------------
# Push images
#----------------------------------------------------------------------------------

.PHONY: push-all-images
push-all-images: service-mesh-hub-image-push

.PHONY: service-mesh-hub-image-push
service-mesh-hub-image-push:
	docker push $(PROJECT_IMAGE):$(VERSION)

#----------------------------------------------------------------------------------
# Helm chart
#----------------------------------------------------------------------------------
HELM_ROOTDIR := install/helm
# Include helm makefile so its targets can be ran from the root of this repo
include install/helm/helm.mk

# Generate Manifests from Helm Chart
.PHONY: chart-gen
chart-gen:
	go run -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) codegen/generate.go -chart

.PHONY: manifest-gen
manifest-gen: install/service-mesh-hub-default.yaml
install/service-mesh-hub-default.yaml: chart-gen
	helm template --include-crds --namespace service-mesh-hub install/helm/service-mesh-hub > $@

#----------------------------------------------------------------------------------
# Test
#----------------------------------------------------------------------------------

# run all tests
# set TEST_PKG to run a specific test package
.PHONY: run-tests
run-tests:
	ginkgo -r -failFast -trace $(GINKGOFLAGS) \
		-ldflags=$(LDFLAGS) \
		-gcflags=$(GCFLAGS) \
		-progress \
		-compilers=4 \
		-skipPackage=$(SKIP_PACKAGES) $(TEST_PKG)

# regen code+manifests, image build+push, and run all tests
# convenience for local testing
.PHONY: test-everything
test-everything: clean-generated-code generated-code manifest-gen run-tests

# TODO(ilackarms): release docs, github assets, and chart
##----------------------------------------------------------------------------------
## Release
##----------------------------------------------------------------------------------
#
#.PHONY: upload-github-release-assets
#upload-github-release-assets: build-cli
#ifeq ($(RELEASE),"true")
#	go run ci/upload_github_release_assets.go
#endif
#
#.PHONY: publish-docs
#publish-docs:
#ifeq ($(RELEASE),"true")
#	make -C docs latest \
#		VERSION=$(VERSION) \
#		RELEASE=$(RELEASE)
#endif

#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

.PHONY: clean
clean:
	rm -rf  _output/ vendor_any/

.PHONY: clean-generated-code
clean-generated-code:
	find pkg -name "*.pb.go" -type f -delete
	find pkg -name "*.hash.go" -type f -delete
	find pkg -name "*.gen.go" -type f -delete
	find pkg -name "*deepcopy.go" -type f -delete
