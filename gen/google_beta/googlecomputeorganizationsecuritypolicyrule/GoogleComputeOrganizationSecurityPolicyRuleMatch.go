package googlecomputeorganizationsecuritypolicyrule


type GoogleComputeOrganizationSecurityPolicyRuleMatch struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_organization_security_policy_rule#config GoogleComputeOrganizationSecurityPolicyRule#config}
	Config *GoogleComputeOrganizationSecurityPolicyRuleMatchConfig `field:"optional" json:"config" yaml:"config"`
	// A description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_organization_security_policy_rule#description GoogleComputeOrganizationSecurityPolicyRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_organization_security_policy_rule#expr GoogleComputeOrganizationSecurityPolicyRule#expr}
	Expr *GoogleComputeOrganizationSecurityPolicyRuleMatchExpr `field:"optional" json:"expr" yaml:"expr"`
	// Preconfigured versioned expression.
	//
	// For organization security policy rules,
	// the only supported type is "SRC_IPS_V1".
	// **NOTE** : 'FIREWALL' type is deprecated. Please use 'google_compute_firewall_policy_rule' resource instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_organization_security_policy_rule#versioned_expr GoogleComputeOrganizationSecurityPolicyRule#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

