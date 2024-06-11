CDKTF_GENERATOR=.bin/cdktf-provider-gen
CDKTF_PACKAGES=$(patsubst %.yml,%,$(wildcard *.yml))
CDKTF_VERSION=0.20.7

.PHONY : all
all: install $(CDKTF_PACKAGES)

.PHONY: install
install:
	go build -o $(CDKTF_GENERATOR) github.com/sourcegraph/cdktf-provider-gen/cmd/cdktf-provider-gen

.PHONY: targets
targets:
	@echo $(CDKTF_PACKAGES)

%: ./%.yml
	$(CDKTF_GENERATOR) --config $@.yml --cdktf-version $(CDKTF_VERSION)
