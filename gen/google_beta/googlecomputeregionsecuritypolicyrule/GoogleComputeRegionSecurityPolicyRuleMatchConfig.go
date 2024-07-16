package googlecomputeregionsecuritypolicyrule


type GoogleComputeRegionSecurityPolicyRuleMatchConfig struct {
	// CIDR IP address range. Maximum number of srcIpRanges allowed is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_security_policy_rule#src_ip_ranges GoogleComputeRegionSecurityPolicyRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
}

