package networkfirewallrulegroup

type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributes struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#destination NetworkfirewallRuleGroup#destination}
	Destination any `field:"optional" json:"destination" yaml:"destination"`
	// destination_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#destination_port NetworkfirewallRuleGroup#destination_port}
	DestinationPort any `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#protocols NetworkfirewallRuleGroup#protocols}.
	Protocols *[]*float64 `field:"optional" json:"protocols" yaml:"protocols"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#source NetworkfirewallRuleGroup#source}
	Source any `field:"optional" json:"source" yaml:"source"`
	// source_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#source_port NetworkfirewallRuleGroup#source_port}
	SourcePort any `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// tcp_flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_rule_group#tcp_flag NetworkfirewallRuleGroup#tcp_flag}
	TcpFlag any `field:"optional" json:"tcpFlag" yaml:"tcpFlag"`
}
