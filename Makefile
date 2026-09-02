CDKTN_GENERATOR=.bin/cdktf-provider-gen
CDKTN_PACKAGES=$(patsubst %.yml,%,$(wildcard *.yml))
CDKTN_VERSION=0.24.0

.PHONY : all
all: install $(CDKTN_PACKAGES)

.PHONY: install
install:
	go build -o $(CDKTN_GENERATOR) github.com/sourcegraph/cdktf-provider-gen/cmd/cdktf-provider-gen

.PHONY: targets
targets:
	@echo $(CDKTN_PACKAGES)

%: ./%.yml
	$(CDKTN_GENERATOR) --config $@.yml --cdktn-version $(CDKTN_VERSION)
