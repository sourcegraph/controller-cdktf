package dataobservemonitorv2


type DataObserveMonitorV2Scheduling struct {
	// interval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#interval DataObserveMonitorV2#interval}
	Interval interface{} `field:"optional" json:"interval" yaml:"interval"`
	// scheduled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#scheduled DataObserveMonitorV2#scheduled}
	Scheduled interface{} `field:"optional" json:"scheduled" yaml:"scheduled"`
	// transform block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#transform DataObserveMonitorV2#transform}
	Transform interface{} `field:"optional" json:"transform" yaml:"transform"`
}

