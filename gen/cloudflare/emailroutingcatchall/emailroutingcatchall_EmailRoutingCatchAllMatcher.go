package emailroutingcatchall


type EmailRoutingCatchAllMatcher struct {
	// Type of matcher. Available values: `all`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_catch_all#type EmailRoutingCatchAll#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

