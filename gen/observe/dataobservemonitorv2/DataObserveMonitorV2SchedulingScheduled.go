package dataobservemonitorv2


type DataObserveMonitorV2SchedulingScheduled struct {
	// A timezone is required to ensure that interpretation of scheduling on the wall-clock is done relative to the desired timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#timezone DataObserveMonitorV2#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

