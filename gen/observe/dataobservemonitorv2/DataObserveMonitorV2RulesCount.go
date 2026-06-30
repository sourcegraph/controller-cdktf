package dataobservemonitorv2

type DataObserveMonitorV2RulesCount struct {
	// compare_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#compare_groups DataObserveMonitorV2#compare_groups}
	CompareGroups any `field:"optional" json:"compareGroups" yaml:"compareGroups"`
	// compare_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#compare_values DataObserveMonitorV2#compare_values}
	CompareValues any `field:"optional" json:"compareValues" yaml:"compareValues"`
}
