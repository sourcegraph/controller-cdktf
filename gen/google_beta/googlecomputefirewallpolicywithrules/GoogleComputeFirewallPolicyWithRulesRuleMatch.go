package googlecomputefirewallpolicywithrules


type GoogleComputeFirewallPolicyWithRulesRuleMatch struct {
	// layer4_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#layer4_config GoogleComputeFirewallPolicyWithRules#layer4_config}
	Layer4Config interface{} `field:"required" json:"layer4Config" yaml:"layer4Config"`
	// Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#dest_address_groups GoogleComputeFirewallPolicyWithRules#dest_address_groups}
	DestAddressGroups *[]*string `field:"optional" json:"destAddressGroups" yaml:"destAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic destination. Maximum number of destination fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#dest_fqdns GoogleComputeFirewallPolicyWithRules#dest_fqdns}
	DestFqdns *[]*string `field:"optional" json:"destFqdns" yaml:"destFqdns"`
	// Destination IP address range in CIDR format. Required for EGRESS rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#dest_ip_ranges GoogleComputeFirewallPolicyWithRules#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// Network scope of the traffic destination. Possible values: ["INTERNET", "INTRA_VPC", "NON_INTERNET", "VPC_NETWORKS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#dest_network_scope GoogleComputeFirewallPolicyWithRules#dest_network_scope}
	DestNetworkScope *string `field:"optional" json:"destNetworkScope" yaml:"destNetworkScope"`
	// Region codes whose IP addresses will be used to match for destination of traffic.
	//
	// Should be specified as 2 letter country code defined as per
	// ISO 3166 alpha-2 country codes. ex."US"
	// Maximum number of destination region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#dest_region_codes GoogleComputeFirewallPolicyWithRules#dest_region_codes}
	DestRegionCodes *[]*string `field:"optional" json:"destRegionCodes" yaml:"destRegionCodes"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#dest_threat_intelligences GoogleComputeFirewallPolicyWithRules#dest_threat_intelligences}
	DestThreatIntelligences *[]*string `field:"optional" json:"destThreatIntelligences" yaml:"destThreatIntelligences"`
	// Address groups which should be matched against the traffic source. Maximum number of source address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#src_address_groups GoogleComputeFirewallPolicyWithRules#src_address_groups}
	SrcAddressGroups *[]*string `field:"optional" json:"srcAddressGroups" yaml:"srcAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic source. Maximum number of source fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#src_fqdns GoogleComputeFirewallPolicyWithRules#src_fqdns}
	SrcFqdns *[]*string `field:"optional" json:"srcFqdns" yaml:"srcFqdns"`
	// Source IP address range in CIDR format. Required for INGRESS rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#src_ip_ranges GoogleComputeFirewallPolicyWithRules#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
	// Networks of the traffic source. It can be either a full or partial url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#src_networks GoogleComputeFirewallPolicyWithRules#src_networks}
	SrcNetworks *[]*string `field:"optional" json:"srcNetworks" yaml:"srcNetworks"`
	// Network scope of the traffic source. Possible values: ["INTERNET", "INTRA_VPC", "NON_INTERNET", "VPC_NETWORKS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#src_network_scope GoogleComputeFirewallPolicyWithRules#src_network_scope}
	SrcNetworkScope *string `field:"optional" json:"srcNetworkScope" yaml:"srcNetworkScope"`
	// Region codes whose IP addresses will be used to match for source of traffic.
	//
	// Should be specified as 2 letter country code defined as per
	// ISO 3166 alpha-2 country codes. ex."US"
	// Maximum number of source region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#src_region_codes GoogleComputeFirewallPolicyWithRules#src_region_codes}
	SrcRegionCodes *[]*string `field:"optional" json:"srcRegionCodes" yaml:"srcRegionCodes"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_firewall_policy_with_rules#src_threat_intelligences GoogleComputeFirewallPolicyWithRules#src_threat_intelligences}
	SrcThreatIntelligences *[]*string `field:"optional" json:"srcThreatIntelligences" yaml:"srcThreatIntelligences"`
}

