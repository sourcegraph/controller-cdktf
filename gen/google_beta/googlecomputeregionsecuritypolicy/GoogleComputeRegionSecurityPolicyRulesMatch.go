package googlecomputeregionsecuritypolicy


type GoogleComputeRegionSecurityPolicyRulesMatch struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_security_policy#config GoogleComputeRegionSecurityPolicy#config}
	Config *GoogleComputeRegionSecurityPolicyRulesMatchConfig `field:"optional" json:"config" yaml:"config"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_security_policy#expr GoogleComputeRegionSecurityPolicy#expr}
	Expr *GoogleComputeRegionSecurityPolicyRulesMatchExpr `field:"optional" json:"expr" yaml:"expr"`
	// Preconfigured versioned expression.
	//
	// If this field is specified, config must also be specified.
	// Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding srcIpRange field in config. Possible values: ["SRC_IPS_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_security_policy#versioned_expr GoogleComputeRegionSecurityPolicy#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

