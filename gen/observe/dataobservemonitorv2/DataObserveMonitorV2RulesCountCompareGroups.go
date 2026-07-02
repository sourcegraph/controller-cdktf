package dataobservemonitorv2

type DataObserveMonitorV2RulesCountCompareGroups struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#column DataObserveMonitorV2#column}
	Column any `field:"optional" json:"column" yaml:"column"`
	// compare_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#compare_values DataObserveMonitorV2#compare_values}
	CompareValues any `field:"optional" json:"compareValues" yaml:"compareValues"`
}
