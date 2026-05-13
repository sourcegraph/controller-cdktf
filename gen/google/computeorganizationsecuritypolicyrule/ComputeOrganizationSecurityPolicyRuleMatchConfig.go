package computeorganizationsecuritypolicyrule


type ComputeOrganizationSecurityPolicyRuleMatchConfig struct {
	// Source IP address range in CIDR format. Required for INGRESS rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_rule#src_ip_ranges ComputeOrganizationSecurityPolicyRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
}

