# Set these to the desired values
PROJECT_NAME=k8s-blueprint-lib
ARTIFACT_ID=k8s-blueprint-operator-crd
APPEND_CRD_SUFFIX=false
VERSION=1.3.0

GOTAG?=1.24.3
MAKEFILES_VERSION=9.10.0

GO_BUILD_FLAGS?=-mod=vendor -a ./...
.DEFAULT_GOAL:=default

PRE_COMPILE = generate-deepcopy
IMAGE_IMPORT_TARGET=image-import
CHECK_VAR_TARGETS=check-all-vars-without-image

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

include build/make/digital-signature.mk
include build/make/k8s-component.mk
include build/make/k8s-crd.mk


default: compile

# Override make target to use k8s-blueprint-lib as label
.PHONY: crd-add-labels
crd-add-labels: $(BINARY_YQ)
	@echo "Adding labels to CRD..."
	@for file in ${HELM_CRD_SOURCE_DIR}/templates/*.yaml ; do \
		$(BINARY_YQ) -i e ".metadata.labels.app = \"ces\"" $${file} ;\
		$(BINARY_YQ) -i e ".metadata.labels.\"app.kubernetes.io/name\" = \"${PROJECT_NAME}\"" $${file} ;\
	done
