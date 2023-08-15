CDKTF_GENERATOR=cdktf-provider-gen
CDKTF_PACKAGES=$(patsubst %.yml,%,$(wildcard *.yml))

.PHONY : all
all: $(CDKTF_PACKAGES)

.PHONY: targets
targets:
	@echo $(CDKTF_PACKAGES)

%: ./%.yml
	$(CDKTF_GENERATOR) --config $@.yml
