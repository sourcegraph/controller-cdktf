package eventhook


type EventHookHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/event_hook#key EventHook#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/event_hook#value EventHook#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

