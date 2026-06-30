package computeregionsecuritypolicy

type ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusion struct {
	// Target WAF rule set to apply the preconfigured WAF exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_security_policy#target_rule_set ComputeRegionSecurityPolicy#target_rule_set}
	TargetRuleSet *string `field:"required" json:"targetRuleSet" yaml:"targetRuleSet"`
	// request_cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_security_policy#request_cookie ComputeRegionSecurityPolicy#request_cookie}
	RequestCookie any `field:"optional" json:"requestCookie" yaml:"requestCookie"`
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_security_policy#request_header ComputeRegionSecurityPolicy#request_header}
	RequestHeader any `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// request_query_param block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_security_policy#request_query_param ComputeRegionSecurityPolicy#request_query_param}
	RequestQueryParam any `field:"optional" json:"requestQueryParam" yaml:"requestQueryParam"`
	// request_uri block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_security_policy#request_uri ComputeRegionSecurityPolicy#request_uri}
	RequestUri any `field:"optional" json:"requestUri" yaml:"requestUri"`
	// A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion.
	//
	// If omitted, it refers to all the rule IDs under the WAF rule set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_security_policy#target_rule_ids ComputeRegionSecurityPolicy#target_rule_ids}
	TargetRuleIds *[]*string `field:"optional" json:"targetRuleIds" yaml:"targetRuleIds"`
}
