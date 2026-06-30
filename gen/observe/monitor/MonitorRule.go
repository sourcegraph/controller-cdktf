package monitor

type MonitorRule struct {
	// change block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#change Monitor#change}
	Change *MonitorRuleChange `field:"optional" json:"change" yaml:"change"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#count Monitor#count}
	Count *MonitorRuleCount `field:"optional" json:"count" yaml:"count"`
	// facet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#facet Monitor#facet}
	Facet *MonitorRuleFacet `field:"optional" json:"facet" yaml:"facet"`
	// group_by_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#group_by_group Monitor#group_by_group}
	GroupByGroup any `field:"optional" json:"groupByGroup" yaml:"groupByGroup"`
	// log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#log Monitor#log}
	Log *MonitorRuleLog `field:"optional" json:"log" yaml:"log"`
	// promote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#promote Monitor#promote}
	Promote *MonitorRulePromote `field:"optional" json:"promote" yaml:"promote"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#source_column Monitor#source_column}.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#threshold Monitor#threshold}
	Threshold *MonitorRuleThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}
