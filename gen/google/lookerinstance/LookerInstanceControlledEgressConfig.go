package lookerinstance


type LookerInstanceControlledEgressConfig struct {
	// List of fully qualified domain names to be added to the allowlist for outbound traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/looker_instance#egress_fqdns LookerInstance#egress_fqdns}
	EgressFqdns *[]*string `field:"optional" json:"egressFqdns" yaml:"egressFqdns"`
	// Whether the Looker Marketplace is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/looker_instance#marketplace_enabled LookerInstance#marketplace_enabled}
	MarketplaceEnabled interface{} `field:"optional" json:"marketplaceEnabled" yaml:"marketplaceEnabled"`
}

