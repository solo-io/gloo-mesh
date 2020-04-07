HELM_DIR := install/helm
HELM_ROOTDIR := $(shell pwd)/$(HELM_DIR)

OUTPUT_DIR ?= $(ROOTDIR)/_output
HELM_OUTPUT_DIR ?= $(OUTPUT_DIR)/helm

# Helm chart sources
CHARTS_DIR := $(HELM_ROOTDIR)/charts
PACKAGED_CHARTS_DIR := $(HELM_OUTPUT_DIR)/charts
# list of SMH component directories, must be a subdirectory in install/helm/charts/
CHARTS := $(COMPONENTS) custom-resource-definitions
RELEASE := "true"
ifeq ($(TAGGED_VERSION),)
# TAGGED_VERSION := $(shell git describe --tags)
# This doesn't work in CI, need to find another way...
	TAGGED_VERSION := $(shell git describe --tags --dirty --always)
	RELEASE := "false"
endif
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)

# for a helm source directory in the install/helm/charts/ directory,
# 1. include install/helm/charts/.helmignore
# 2. package the chart into a new directory under $(PACKAGED_CHARTS_DIR)
define package_chart
cp $(HELM_ROOTDIR)/.helmignore $(CHARTS_DIR)/$(1)/;
mkdir -p $(PACKAGED_CHARTS_DIR);
helm package --destination $(PACKAGED_CHARTS_DIR)/$(1) $(CHARTS_DIR)/$(1);
endef

# create the helm repo index.yaml in the packaged directory
define index_chart
helm repo index $(PACKAGED_CHARTS_DIR)/$(1);
endef

define push_chart_to_registry
HELM_EXPERIMENTAL_OCI=1 helm chart save $(CHARTS_DIR)/$(1) gcr.io/service-mesh-hub/$(1):$(VERSION);
HELM_EXPERIMENTAL_OCI=1 helm chart push gcr.io/service-mesh-hub/$(1):$(VERSION);
endef

# invoked once per dependency to be copied
# $(1) is the component to be packaged, $(2) is the dependency
define copy_as_dependency
mkdir -p $(CHARTS_DIR)/$(1)/charts/$(2);
cp -r $(CHARTS_DIR)/$(2) $(CHARTS_DIR)/$(1)/charts/;
endef

# make a copy of */Chart-template.yaml and */values-template.yaml to */Chart.yaml and */values.yaml, with version injection
.PHONY: set-version
set-version:
	find $(CHARTS_DIR) -type f -name "Chart-template.yaml" | while read f; do sed -e 's/%version%/'$(VERSION)'/' $$f > $$(dirname $$f)/Chart.yaml; done
	find $(CHARTS_DIR) -type f -name "values-template.yaml" | while read f; do sed -e 's/%version%/'$(VERSION)'/' $$f > $$(dirname $$f)/values.yaml; done

# for each component in $(CHARTS), whose name must be a subdirectory in install/helm/charts/,
# package the helm chart
.PHONY: package-component-charts
package-component-charts:
	$(foreach chart,$(CHARTS),$(call package_chart,$(chart)))

# generate the helm repo index.yaml for SMH components
.PHONY: index-component-charts
index-component-charts:
	$(foreach chart,$(CHARTS),$(call index_chart,$(chart)))

# package, index, and upload Helm charts SMH components
.PHONY: package-index-components-helm
package-index-components-helm: set-version package-component-charts index-component-charts

#----------------------------------------------------------------------------------
# management-plane
#----------------------------------------------------------------------------------

# copy component dependencies into app chart
.PHONY: copy-dependencies-mgmt-plane
copy-dependencies-mgmt-plane:
	$(foreach chart,$(CHARTS),$(call copy_as_dependency,management-plane,$(chart)))

# fetch all management-plane Helm dependencies (SMH component charts), then package management-plane
.PHONY: package-mgmt-plane-chart
package-mgmt-plane-chart:
	$(call package_chart,management-plane)

# generate the helm repo index.yaml for SMH app
.PHONY: index-mgmt-plane-chart
index-mgmt-plane-chart:
	$(call index_chart,management-plane)

# package, index, and upload Helm chart for the SMH app
package-index-mgmt-plane-helm: package-index-components-helm copy-dependencies-mgmt-plane package-mgmt-plane-chart index-mgmt-plane-chart

#----------------------------------------------------------------------------------
# csr-agent
#----------------------------------------------------------------------------------

# copy component dependencies into app chart
.PHONY: copy-dependencies-csr-agent
copy-dependencies-csr-agent:
	$(call copy_as_dependency,csr-agent,custom-resource-definitions)

# fetch all SMH Helm dependencies (SMH component charts), then package SMH
.PHONY: package-csr-agent-chart
package-csr-agent-chart:
	$(call package_chart,csr-agent)

# generate the helm repo index.yaml for SMH app
.PHONY: index-csr-agent-chart
index-csr-agent-chart:
	$(call index_chart,csr-agent)

# package, index, and upload Helm chart for the SMH csr-agent
package-index-csr-agent-helm: package-index-components-helm copy-dependencies-csr-agent package-csr-agent-chart index-csr-agent-chart

#----------------------------------------------------------------------------------
# Build Targets
#----------------------------------------------------------------------------------

# upload the new Helm package and index.yaml
.PHONY: save-helm
save-helm:
ifeq ($(RELEASE),"true")
	gsutil -m rsync -r -d $(HELM_OUTPUT_DIR)/charts gs://service-mesh-hub/
else
	echo "Not a release, skipping uploading to GCS."
endif

# must be executed during build prior to indexing helm repo to maintain prior versions in repo
.PHONY: fetch-helm
fetch-helm:
	mkdir -p $(HELM_OUTPUT_DIR)/charts
	gsutil -m rsync -r gs://service-mesh-hub/ $(HELM_OUTPUT_DIR)/charts

# upload Helm chart to GCR as an OCI image
.PHONY: push-chart-to-registry
push-chart-to-registry:
	mkdir -p $(HELM_REPOSITORY_CACHE)
	cp $(DOCKER_CONFIG)/config.json $(HELM_REPOSITORY_CACHE)/config.json
	$(foreach chart,$(CHARTS),$(call push_chart_to_registry,$(chart)))
	$(call push_chart_to_registry,management-plane)

# package, index, and upload Helm charts for SMH components, then the management-plane, then the csr-agent
.PHONY: release-helm
release-helm: package-index-mgmt-plane-helm package-index-csr-agent-helm save-helm

.PHONY: helm-clean
helm-clean:
	rm -rf $(HELM_OUTPUT_DIR)
	rm -rf $(CHARTS_DIR)/management-plane/charts
	rm -rf $(CHARTS_DIR)/management-plane/Chart.lock
	rm -rf $(CHARTS_DIR)/csr-agent/charts
	rm -rf $(CHARTS_DIR)/csr-agent/Chart.lock
	find $(CHARTS_DIR) -type f -name "Chart.yaml" -exec rm {} \;
	find $(CHARTS_DIR) -type f -name "values.yaml" -exec rm {} \;
