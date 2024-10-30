package inlinehook


type InlineHookHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/inline_hook#key InlineHook#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/inline_hook#value InlineHook#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

