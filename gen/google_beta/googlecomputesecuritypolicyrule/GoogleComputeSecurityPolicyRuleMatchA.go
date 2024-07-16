package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRuleMatchA struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy_rule#config GoogleComputeSecurityPolicyRuleA#config}
	Config *GoogleComputeSecurityPolicyRuleMatchConfigA `field:"optional" json:"config" yaml:"config"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy_rule#expr GoogleComputeSecurityPolicyRuleA#expr}
	Expr *GoogleComputeSecurityPolicyRuleMatchExprA `field:"optional" json:"expr" yaml:"expr"`
	// expr_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy_rule#expr_options GoogleComputeSecurityPolicyRuleA#expr_options}
	ExprOptions *GoogleComputeSecurityPolicyRuleMatchExprOptionsA `field:"optional" json:"exprOptions" yaml:"exprOptions"`
	// Preconfigured versioned expression.
	//
	// If this field is specified, config must also be specified.
	// Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding srcIpRange field in config. Possible values: ["SRC_IPS_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy_rule#versioned_expr GoogleComputeSecurityPolicyRuleA#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

