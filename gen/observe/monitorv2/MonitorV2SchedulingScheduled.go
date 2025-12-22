package monitorv2


type MonitorV2SchedulingScheduled struct {
	// A timezone is required to ensure that interpretation of scheduling on the wall-clock is done relative to the desired timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#timezone MonitorV2#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
	// If specified, the raw cron is a crontab configuration to use to drive the scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#raw_cron MonitorV2#raw_cron}
	RawCron *string `field:"optional" json:"rawCron" yaml:"rawCron"`
}

