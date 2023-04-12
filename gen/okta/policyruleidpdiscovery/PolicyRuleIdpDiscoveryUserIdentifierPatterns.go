package policyruleidpdiscovery


type PolicyRuleIdpDiscoveryUserIdentifierPatterns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_idp_discovery#match_type PolicyRuleIdpDiscovery#match_type}.
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_idp_discovery#value PolicyRuleIdpDiscovery#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

