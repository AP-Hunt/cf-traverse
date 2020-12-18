VERSION_MAJOR := $(shell semver-cli get major $$(cat ./version.txt))
VERSION_MINOR := $(shell semver-cli get minor $$(cat ./version.txt))
VERSION_PATCH := $(shell semver-cli get patch $$(cat ./version.txt))

build:
	go build \
      -o bin/cf-traverse \
      -ldflags="-X 'github.com/AP-Hunt/cf-traverse/version.MAJOR_VERSION=${VERSION_MAJOR}' -X 'github.com/AP-Hunt/cf-traverse/version.MINOR_VERSION=${VERSION_MINOR}' -X 'github.com/AP-Hunt/cf-traverse/version.PATCH_VERSION=${VERSION_PATCH}'" \
      .

install: build
	cf install-plugin -f ./bin/cf-traverse
