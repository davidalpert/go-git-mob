PROJECTNAME=git-mob

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# go versioning flags
ifndef VERSION
	VERSION=$(shell sbot get version)
endif

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)

# ---------------------- targets -------------------------------------

.PHONY: default
default: help

.PHONY: cit
cit: clean vale build test-unit test-features ## clean build and test-all (except @wip features)

.PHONY: version
version: ## show current version
	echo ${VERSION}

.PHONY: clean
clean: ## clean build output
	rm -rf ./bin

.PHONY: vale
vale: ## run linting rules against markdown files
ifeq ("$(GITHUB_ACTIONS)","true")
	echo "GITHUB_ACTIONS is true; assuming vale has been run by a previous step"
else
	vale README.md CONTRIBUTING.md # we don't valedate CHANGELOG.md as that reflects historical commit summaries
endif

# SRC_FILES = $(shell find . -type f -name '*.go' -not -path "./internal/version/detail.go")
# SRC_FILES += Makefile

# this needs to be a PHONY target so that it runs all the time, otherwise changing the VERSION is not enough to trigger an update to the version detail
.PHONY: ./internal/version/detail.go
./internal/version/detail.go:
	VERSION=$(VERSION) go run ./.tools/version_gen.go $(PROJECTNAME)

.PHONY: gen
gen: ## invoke go generate
	echo "invoking go generate (with VERSION=$(VERSION))"
	@CGO_ENABLED=1 VERSION=$(VERSION) go generate ./...

.PHONY: build
build: clean ./internal/version/detail.go ## build for current platform
	mkdir -p ./bin
	go build -o ./bin/git-mob main.go

.PHONY: build-all
build-all: clean ./internal/version/detail.go gen ## build for all platforms
	GOOS=darwin GOARCH=arm64 mkdir -p bin/darwin-arm64
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64/git-mob main.go
	GOOS=darwin GOARCH=amd64 mkdir -p bin/darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64/git-mob main.go
	GOOS=linux GOARCH=amd64 mkdir -p bin/linux-amd64
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/git-mob main.go

.PHONY: install
install: build ## build and install locally into GOPATH
	echo "writing: ${GOPATH}/bin/git-mob"
	cp ./bin/git-mob ${GOPATH}/bin
	${GOPATH}/bin/git-mob install

.PHONY: implode
implode: uninstall ## uninstall locally from GOPATH

.PHONY: uninstall
uninstall: ## uninstall locally from GOPATH
# ifneq ("$(wildcard $(${GOPATH}/bin/git-mob))","")
ifneq ("$(shell which git-mob)","")
	${GOPATH}/bin/git-mob uninstall
else
	echo "git-mob not found in GOPATH"
endif

.PHONY: test
test: test-unit test-features ## run all tests (unit and integration)

.PHONY: test-unit
test-unit: ./internal/version/detail.go ## run unit tests
	go test -v ./...

.PHONY: test-features
test-features: Gemfile.lock build ## Run cucumber/aruba backend features
	bundle exec cucumber --publish-quiet --tags 'not @wip' --tags 'not @ignore'

.PHONY: test-features-wip
test-features-wip: Gemfile.lock build ## Run cucumber/aruba backend features
	bundle exec cucumber --publish-quiet --tags '@wip' --tags 'not @ignore'

.PHONY: list-ignored
list-ignored: ## list ignored specs
	bundle exec cucumber --publish-quiet --tags '@ignore' --dry-run

test-action-push: ## Test github actions with event 'push'
	# https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#webhook-payload-example-37
	act push --container-architecture linux/amd64 --eventpath .local/push-tags-payload.json

.PHONY: depgraph
depgraph: ## create a dotgraph visualizing package dependencies
	 godepgraph -s -novendor -p github.com/go-git/go-git,github.com/spf13,github.com/go-xmlfmt,github.com/gocarina/gocsv,github.com/olekukonko/tablewriter,golang.org,gopkg.in $(shell cat go.mod | grep module | sed -E 's/module //') | dot -T svg > dependencygraph.svg

.PHONY: deploy
deploy: test build-all ## deploy binaries
	$(if $(shell which ./deploy.sh),./deploy.sh,$(error "./deploy.sh not found"))

.PHONY: deploy-local
deploy-local: test build ## deploy binaries locally (for testing)
	$(if $(shell which ./deploy-local.sh),./deploy-local.sh,$(error "./deploy-local.sh not found"))

.PHONY: doctor
doctor: ## run doctor.sh to sort out development dependencies
	./.tools/doctor.sh

.PHONY: changelog
changelog: ## Generate/update CHANGELOG.md
	git-chglog --output CHANGELOG.md

.PHONY: preview-release-notes
preview-release-notes: ## preview release notes (generates RELEASE_NOTES.md)
	git-chglog --output RELEASE_NOTES.md --template .chglog/RELEASE_NOTES.tpl.md "v$(shell sbot get version)"

.PHONY: preview-release
preview-release: preview-release-notes ## preview release (using goreleaser --snapshot)
	goreleaser release --snapshot --rm-dist --release-notes RELEASE_NOTES.md

eq = $(and $(findstring $(1),$(2)),$(findstring $(2),$(1)))

.PHONY: tag-release
tag-release:
	$(if $(call eq,0,$(shell git diff-files --quiet; echo $$?)),, \
		$(error There are unstaged changes; clean your working directory before releasing.) \
	)
	$(if $(call eq,0,$(shell git diff-index --quiet --cached HEAD --; echo $$?)),, \
		$(error There are uncomitted changes; clean your working directory before releasing.) \
	)
	$(eval next_version := $(shell sbot predict version --mode ${BUMP_TYPE}))
	# echo "Current Version: ${VERSION}"
	# echo "   Next Version: ${next_version}"
	make build
ifdef FAST
	$(MAKE) test-unit VERSION=$(next_version)
else
	$(MAKE) cit VERSION=$(next_version)
endif
	git add -f internal/version/detail.go
	git-chglog --next-tag v$(next_version) --output CHANGELOG.md
	git add -f CHANGELOG.md
	git commit --message "release notes for v$(next_version)"
	sbot release version --mode ${BUMP_TYPE}
	git show --no-patch --format=short v$(next_version)

SEMVER_TYPES := major minor patch
BUMP_TARGETS := $(addprefix release-,$(SEMVER_TYPES))
.PHONY: $(BUMP_TARGETS)
$(BUMP_TARGETS): ## bump version
	$(eval BUMP_TYPE := $(strip $(word 2,$(subst -, ,$@))))
	$(MAKE) tag-release BUMP_TYPE=$(BUMP_TYPE)

.PHONY: help
help: Makefile
	@echo
	@echo " ${PROJECTNAME} ${VERSION} - available targets:"
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	printf "\033[36m%-30s\033[0m %s\n" '----------' '------------------'
	@echo $(BUMP_TARGETS) | tr ' ' '\n' | sort | sed -E 's/((.+)\-(.+))/\1: ## \2 \3 version/' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
