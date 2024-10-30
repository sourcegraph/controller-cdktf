package policyrulesignon


type PolicyRuleSignonFactorSequence struct {
	// Type of a Factor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_signon#primary_criteria_factor_type PolicyRuleSignon#primary_criteria_factor_type}
	PrimaryCriteriaFactorType *string `field:"required" json:"primaryCriteriaFactorType" yaml:"primaryCriteriaFactorType"`
	// Factor provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_signon#primary_criteria_provider PolicyRuleSignon#primary_criteria_provider}
	PrimaryCriteriaProvider *string `field:"required" json:"primaryCriteriaProvider" yaml:"primaryCriteriaProvider"`
	// secondary_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_signon#secondary_criteria PolicyRuleSignon#secondary_criteria}
	SecondaryCriteria interface{} `field:"optional" json:"secondaryCriteria" yaml:"secondaryCriteria"`
}

