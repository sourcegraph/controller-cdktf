package networkpolicy


type NetworkPolicySpecIngressFromIpBlock struct {
	// CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/network_policy#cidr NetworkPolicy#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64" Except values will be rejected if they are outside the CIDR range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/network_policy#except NetworkPolicy#except}
	Except *[]*string `field:"optional" json:"except" yaml:"except"`
}

