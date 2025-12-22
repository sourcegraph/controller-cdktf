package monitorv2


type MonitorV2ActionsConditionsCompareTerms struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#column MonitorV2#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// comparison block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#comparison MonitorV2#comparison}
	Comparison interface{} `field:"required" json:"comparison" yaml:"comparison"`
}

