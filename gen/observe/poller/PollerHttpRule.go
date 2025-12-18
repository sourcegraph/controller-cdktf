package poller


type PollerHttpRule struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#match Poller#match}
	Match *PollerHttpRuleMatch `field:"required" json:"match" yaml:"match"`
	// decoder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#decoder Poller#decoder}
	Decoder *PollerHttpRuleDecoder `field:"optional" json:"decoder" yaml:"decoder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#follow Poller#follow}.
	Follow *string `field:"optional" json:"follow" yaml:"follow"`
}

