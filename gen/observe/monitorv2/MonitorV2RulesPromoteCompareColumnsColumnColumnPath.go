package monitorv2


type MonitorV2RulesPromoteCompareColumnsColumnColumnPath struct {
	// The name of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#name MonitorV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path of the path, if the name refers to a column with a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#path MonitorV2#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

