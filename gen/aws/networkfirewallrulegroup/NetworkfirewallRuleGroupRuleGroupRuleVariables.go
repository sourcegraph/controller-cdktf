package networkfirewallrulegroup

type NetworkfirewallRuleGroupRuleGroupRuleVariables struct {
	// ip_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#ip_sets NetworkfirewallRuleGroup#ip_sets}
	IpSets any `field:"optional" json:"ipSets" yaml:"ipSets"`
	// port_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#port_sets NetworkfirewallRuleGroup#port_sets}
	PortSets any `field:"optional" json:"portSets" yaml:"portSets"`
}
