package dataobservemonitor


type DataObserveMonitorRuleThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#compare_function DataObserveMonitor#compare_function}.
	CompareFunction *string `field:"optional" json:"compareFunction" yaml:"compareFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#compare_values DataObserveMonitor#compare_values}.
	CompareValues *[]*float64 `field:"optional" json:"compareValues" yaml:"compareValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#lookback_time DataObserveMonitor#lookback_time}.
	LookbackTime *string `field:"optional" json:"lookbackTime" yaml:"lookbackTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#threshold_agg_function DataObserveMonitor#threshold_agg_function}.
	ThresholdAggFunction *string `field:"optional" json:"thresholdAggFunction" yaml:"thresholdAggFunction"`
}

