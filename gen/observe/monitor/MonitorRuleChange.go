package monitor


type MonitorRuleChange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#baseline_time Monitor#baseline_time}.
	BaselineTime *string `field:"required" json:"baselineTime" yaml:"baselineTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#compare_function Monitor#compare_function}.
	CompareFunction *string `field:"required" json:"compareFunction" yaml:"compareFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#lookback_time Monitor#lookback_time}.
	LookbackTime *string `field:"required" json:"lookbackTime" yaml:"lookbackTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#aggregate_function Monitor#aggregate_function}.
	AggregateFunction *string `field:"optional" json:"aggregateFunction" yaml:"aggregateFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#change_type Monitor#change_type}.
	ChangeType *string `field:"optional" json:"changeType" yaml:"changeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#compare_value Monitor#compare_value}.
	CompareValue *float64 `field:"optional" json:"compareValue" yaml:"compareValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#compare_values Monitor#compare_values}.
	CompareValues *[]*float64 `field:"optional" json:"compareValues" yaml:"compareValues"`
}

