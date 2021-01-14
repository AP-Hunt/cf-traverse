VERSION_MAJOR := $(shell semver-cli get major $$(cat ./version.txt))
VERSION_MINOR := $(shell semver-cli get minor $$(cat ./version.txt))
VERSION_PATCH := $(shell semver-cli get patch $$(cat ./version.txt))
XCOMPILE_PLATS := darwin/amd64 linux/386 linux/amd64 windows/386 windows/amd64

build:
	go build \
      -o bin/cf-traverse \
      -ldflags="-X 'github.com/AP-Hunt/cf-traverse/version.MAJOR_VERSION=${VERSION_MAJOR}' -X 'github.com/AP-Hunt/cf-traverse/version.MINOR_VERSION=${VERSION_MINOR}' -X 'github.com/AP-Hunt/cf-traverse/version.PATCH_VERSION=${VERSION_PATCH}'" \
      .

install: build
	cf install-plugin -f ./bin/cf-traverse

release:
	CGO_ENABLED=0 \
	gox \
      -output "bin/release/cf-traverse-${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}-{{.OS}}-{{.Arch}}" \
      -ldflags="-X 'github.com/AP-Hunt/cf-traverse/version.MAJOR_VERSION=${VERSION_MAJOR}' -X 'github.com/AP-Hunt/cf-traverse/version.MINOR_VERSION=${VERSION_MINOR}' -X 'github.com/AP-Hunt/cf-traverse/version.PATCH_VERSION=${VERSION_PATCH}'" \
      -osarch="${XCOMPILE_PLATS}" \
      . && \
      cd bin/release && \
      tar czf "../../cf-traverse-${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}.tar.gz" *
