package monitorv2


type MonitorV2RulesCount struct {
	// compare_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#compare_values MonitorV2#compare_values}
	CompareValues interface{} `field:"required" json:"compareValues" yaml:"compareValues"`
	// compare_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#compare_groups MonitorV2#compare_groups}
	CompareGroups interface{} `field:"optional" json:"compareGroups" yaml:"compareGroups"`
}

