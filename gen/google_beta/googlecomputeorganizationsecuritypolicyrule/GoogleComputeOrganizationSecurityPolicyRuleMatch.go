package googlecomputeorganizationsecuritypolicyrule


type GoogleComputeOrganizationSecurityPolicyRuleMatch struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#config GoogleComputeOrganizationSecurityPolicyRule#config}
	Config *GoogleComputeOrganizationSecurityPolicyRuleMatchConfig `field:"required" json:"config" yaml:"config"`
	// A description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#description GoogleComputeOrganizationSecurityPolicyRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Preconfigured versioned expression. For organization security policy rules, the only supported type is "FIREWALL". Default value: "FIREWALL" Possible values: ["FIREWALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#versioned_expr GoogleComputeOrganizationSecurityPolicyRule#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

