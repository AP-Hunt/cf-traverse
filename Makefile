VERSION_MAJOR := $(shell semver-cli get major $$(cat ./version.txt))
VERSION_MINOR := $(shell semver-cli get minor $$(cat ./version.txt))
VERSION_PATCH := $(shell semver-cli get patch $$(cat ./version.txt))
GO_OS := $(shell go env GOOS)
GO_ARCH := $(shell go env GOARCH)

build:
	go build \
      -o bin/cf-traverse \
      -ldflags="-X 'github.com/AP-Hunt/cf-traverse/version.MAJOR_VERSION=${VERSION_MAJOR}' -X 'github.com/AP-Hunt/cf-traverse/version.MINOR_VERSION=${VERSION_MINOR}' -X 'github.com/AP-Hunt/cf-traverse/version.PATCH_VERSION=${VERSION_PATCH}'" \
      .

install: build
	cf install-plugin -f ./bin/cf-traverse

release: 
	go build \
      -o "bin/release/cf-traverse-${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}-${GO_OS}-${GO_ARCH}" \
      -ldflags="-X 'github.com/AP-Hunt/cf-traverse/version.MAJOR_VERSION=${VERSION_MAJOR}' -X 'github.com/AP-Hunt/cf-traverse/version.MINOR_VERSION=${VERSION_MINOR}' -X 'github.com/AP-Hunt/cf-traverse/version.PATCH_VERSION=${VERSION_PATCH}'" \
      . && \
      cd bin/release && \
      tar czf "../../cf-traverse-${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}.tar.gz" *
