package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRuleMatchExpr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#expression GoogleComputeSecurityPolicy#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

