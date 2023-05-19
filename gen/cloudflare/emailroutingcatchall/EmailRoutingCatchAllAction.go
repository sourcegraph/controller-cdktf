package emailroutingcatchall


type EmailRoutingCatchAllAction struct {
	// Type of supported action. Available values: `drop`, `forward`, `worker`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/email_routing_catch_all#type EmailRoutingCatchAll#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list with items in the following form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/email_routing_catch_all#value EmailRoutingCatchAll#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

