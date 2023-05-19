package signonpolicyrule


type SignonPolicyRuleFactorSequenceSecondaryCriteria struct {
	// Type of a Factor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/signon_policy_rule#factor_type SignonPolicyRule#factor_type}
	FactorType *string `field:"required" json:"factorType" yaml:"factorType"`
	// Factor provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/signon_policy_rule#provider SignonPolicyRule#provider}
	Provider *string `field:"required" json:"provider" yaml:"provider"`
}

