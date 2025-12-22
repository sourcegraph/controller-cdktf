package monitorv2


type MonitorV2RulesPromoteCompareColumns struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#column MonitorV2#column}
	Column *MonitorV2RulesPromoteCompareColumnsColumn `field:"required" json:"column" yaml:"column"`
	// compare_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#compare_values MonitorV2#compare_values}
	CompareValues interface{} `field:"required" json:"compareValues" yaml:"compareValues"`
}

