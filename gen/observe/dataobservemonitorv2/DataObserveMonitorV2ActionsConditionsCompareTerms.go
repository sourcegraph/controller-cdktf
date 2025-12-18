package dataobservemonitorv2


type DataObserveMonitorV2ActionsConditionsCompareTerms struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#column DataObserveMonitorV2#column}
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// comparison block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#comparison DataObserveMonitorV2#comparison}
	Comparison interface{} `field:"optional" json:"comparison" yaml:"comparison"`
}

