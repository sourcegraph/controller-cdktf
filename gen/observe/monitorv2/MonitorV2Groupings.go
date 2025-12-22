package monitorv2


type MonitorV2Groupings struct {
	// column_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#column_path MonitorV2#column_path}
	ColumnPath *MonitorV2GroupingsColumnPath `field:"optional" json:"columnPath" yaml:"columnPath"`
	// link_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#link_column MonitorV2#link_column}
	LinkColumn *MonitorV2GroupingsLinkColumn `field:"optional" json:"linkColumn" yaml:"linkColumn"`
}

