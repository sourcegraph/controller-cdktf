package inlinehook


type InlineHookHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#key InlineHook#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#value InlineHook#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

