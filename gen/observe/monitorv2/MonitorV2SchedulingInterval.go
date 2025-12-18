package monitorv2


type MonitorV2SchedulingInterval struct {
	// How often the monitor should attempt to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#interval MonitorV2#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// A maximum +/- to apply to the interval to avoid things like harmonics and work stacking up in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#randomize MonitorV2#randomize}
	Randomize *string `field:"required" json:"randomize" yaml:"randomize"`
}

