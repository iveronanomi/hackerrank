.PHONY: test/all test/all_domain test/contest

GOPATH:=$(lastword $(subst :, ,$(GOPATH)))
PROJ_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GLIDE_PATH:=$(GOPATH)/src/github.com/Masterminds/glide
GLIDE_VERSION:=0.10.2
GLIDE_BIN:=$(GLIDE_PATH)/glide-$(GLIDE_VERSION)

GO=GO15VENDOREXPERIMENT=1 go
NOVENDOR_DIRS=$$(GO15VENDOREXPERIMENT=1 $(GLIDE_BIN) novendor)


# dirty hack to call "glide install" only when "glide.yaml" was changed
# not for external call

test/all:
	$(info Starting all tests...)
	@cd $(PROJ_PATH) && ${GO} test $(extra) -tags="all" ${NOVENDOR_DIRS}

test/all_domains:
	$(info Starting handlers (integration) tests...)
	@cd $(PROJ_PATH) && ${GO} test $(extra) -tags="all" ./all_domains/...

test/contest:
	$(info Starting handlers (integration) tests...)
	@cd $(PROJ_PATH) && ${GO} test $(extra) -tags="all" ./contest/...