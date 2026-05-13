package googlelookerinstance


type GoogleLookerInstanceControlledEgressConfig struct {
	// List of fully qualified domain names to be added to the allowlist for outbound traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_looker_instance#egress_fqdns GoogleLookerInstance#egress_fqdns}
	EgressFqdns *[]*string `field:"optional" json:"egressFqdns" yaml:"egressFqdns"`
	// Whether the Looker Marketplace is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_looker_instance#marketplace_enabled GoogleLookerInstance#marketplace_enabled}
	MarketplaceEnabled interface{} `field:"optional" json:"marketplaceEnabled" yaml:"marketplaceEnabled"`
}

