package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRuleVariables struct {
	// ip_sets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#ip_sets NetworkfirewallRuleGroup#ip_sets}
	IpSets interface{} `field:"optional" json:"ipSets" yaml:"ipSets"`
	// port_sets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#port_sets NetworkfirewallRuleGroup#port_sets}
	PortSets interface{} `field:"optional" json:"portSets" yaml:"portSets"`
}

