package googlecomputeorganizationsecuritypolicyrule


type GoogleComputeOrganizationSecurityPolicyRuleMatchExpr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_organization_security_policy_rule#expression GoogleComputeOrganizationSecurityPolicyRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

