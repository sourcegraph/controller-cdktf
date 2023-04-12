package signonpolicyrule


type SignonPolicyRuleFactorSequence struct {
	// Type of a Factor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy_rule#primary_criteria_factor_type SignonPolicyRule#primary_criteria_factor_type}
	PrimaryCriteriaFactorType *string `field:"required" json:"primaryCriteriaFactorType" yaml:"primaryCriteriaFactorType"`
	// Factor provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy_rule#primary_criteria_provider SignonPolicyRule#primary_criteria_provider}
	PrimaryCriteriaProvider *string `field:"required" json:"primaryCriteriaProvider" yaml:"primaryCriteriaProvider"`
	// secondary_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy_rule#secondary_criteria SignonPolicyRule#secondary_criteria}
	SecondaryCriteria interface{} `field:"optional" json:"secondaryCriteria" yaml:"secondaryCriteria"`
}

