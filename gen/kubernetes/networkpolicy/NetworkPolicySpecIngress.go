package networkpolicy

type NetworkPolicySpecIngress struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/network_policy#from NetworkPolicy#from}
	From any `field:"optional" json:"from" yaml:"from"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/network_policy#ports NetworkPolicy#ports}
	Ports any `field:"optional" json:"ports" yaml:"ports"`
}
