# Set these to the desired values
ARTIFACT_ID=k8s-blueprint-lib
VERSION=0.1.0

GOTAG?=1.23.4
MAKEFILES_VERSION=9.5.2

GO_BUILD_FLAGS?=-mod=vendor -a ./...
.DEFAULT_GOAL:=default

include build/make/variables.mk
INTEGRATION_TEST_NAME_PATTERN=.*_inttest$$

include build/make/self-update.mk
include build/make/dependencies-gomod.mk
include build/make/build.mk
include build/make/test-common.mk
include build/make/test-integration.mk
include build/make/test-unit.mk
include build/make/static-analysis.mk
include build/make/clean.mk
include build/make/release.mk
include build/make/mocks.mk


default: compile
