package monitor


type MonitorRuleGroupByGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#columns Monitor#columns}.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#group_name Monitor#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
}

