package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRuleMatch struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_security_policy#config GoogleComputeSecurityPolicy#config}
	Config *GoogleComputeSecurityPolicyRuleMatchConfig `field:"optional" json:"config" yaml:"config"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_security_policy#expr GoogleComputeSecurityPolicy#expr}
	Expr *GoogleComputeSecurityPolicyRuleMatchExpr `field:"optional" json:"expr" yaml:"expr"`
	// Predefined rule expression.
	//
	// If this field is specified, config must also be specified. Available options:   SRC_IPS_V1: Must specify the corresponding src_ip_ranges field in config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_security_policy#versioned_expr GoogleComputeSecurityPolicy#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

