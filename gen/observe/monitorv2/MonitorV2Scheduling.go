package monitorv2


type MonitorV2Scheduling struct {
	// interval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#interval MonitorV2#interval}
	Interval *MonitorV2SchedulingInterval `field:"optional" json:"interval" yaml:"interval"`
	// scheduled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#scheduled MonitorV2#scheduled}
	Scheduled *MonitorV2SchedulingScheduled `field:"optional" json:"scheduled" yaml:"scheduled"`
	// transform block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#transform MonitorV2#transform}
	Transform *MonitorV2SchedulingTransform `field:"optional" json:"transform" yaml:"transform"`
}

