package monitor


type MonitorRuleThreshold struct {
	// Comparison function used to compare the query result against the compare_values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#compare_function Monitor#compare_function}
	CompareFunction *string `field:"required" json:"compareFunction" yaml:"compareFunction"`
	// Amount of time to evaluate query and comparison over.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#lookback_time Monitor#lookback_time}
	LookbackTime *string `field:"required" json:"lookbackTime" yaml:"lookbackTime"`
	// Value(s) to compare the query(ies) against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#compare_values Monitor#compare_values}
	CompareValues *[]*float64 `field:"optional" json:"compareValues" yaml:"compareValues"`
	// Function used to aggregate threshold events to determine when an alert should be triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#threshold_agg_function Monitor#threshold_agg_function}
	ThresholdAggFunction *string `field:"optional" json:"thresholdAggFunction" yaml:"thresholdAggFunction"`
}

