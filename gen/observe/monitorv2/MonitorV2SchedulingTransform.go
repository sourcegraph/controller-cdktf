package monitorv2


type MonitorV2SchedulingTransform struct {
	// The freshness goal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#freshness_goal MonitorV2#freshness_goal}
	FreshnessGoal *string `field:"required" json:"freshnessGoal" yaml:"freshnessGoal"`
}

