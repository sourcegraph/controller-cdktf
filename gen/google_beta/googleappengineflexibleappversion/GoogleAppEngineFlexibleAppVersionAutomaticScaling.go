package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionAutomaticScaling struct {
	// cpu_utilization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#cpu_utilization GoogleAppEngineFlexibleAppVersion#cpu_utilization}
	CpuUtilization *GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilization `field:"required" json:"cpuUtilization" yaml:"cpuUtilization"`
	// The time period that the Autoscaler should wait before it starts collecting information from a new instance.
	//
	// This prevents the autoscaler from collecting information when the instance is initializing,
	// during which the collected usage would not be reliable. Default: 120s
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#cool_down_period GoogleAppEngineFlexibleAppVersion#cool_down_period}
	CoolDownPeriod *string `field:"optional" json:"coolDownPeriod" yaml:"coolDownPeriod"`
	// disk_utilization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#disk_utilization GoogleAppEngineFlexibleAppVersion#disk_utilization}
	DiskUtilization *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization `field:"optional" json:"diskUtilization" yaml:"diskUtilization"`
	// Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
	//
	// Defaults to a runtime-specific value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#max_concurrent_requests GoogleAppEngineFlexibleAppVersion#max_concurrent_requests}
	MaxConcurrentRequests *float64 `field:"optional" json:"maxConcurrentRequests" yaml:"maxConcurrentRequests"`
	// Maximum number of idle instances that should be maintained for this version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#max_idle_instances GoogleAppEngineFlexibleAppVersion#max_idle_instances}
	MaxIdleInstances *float64 `field:"optional" json:"maxIdleInstances" yaml:"maxIdleInstances"`
	// Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#max_pending_latency GoogleAppEngineFlexibleAppVersion#max_pending_latency}
	MaxPendingLatency *string `field:"optional" json:"maxPendingLatency" yaml:"maxPendingLatency"`
	// Maximum number of instances that should be started to handle requests for this version. Default: 20.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#max_total_instances GoogleAppEngineFlexibleAppVersion#max_total_instances}
	MaxTotalInstances *float64 `field:"optional" json:"maxTotalInstances" yaml:"maxTotalInstances"`
	// Minimum number of idle instances that should be maintained for this version.
	//
	// Only applicable for the default version of a service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#min_idle_instances GoogleAppEngineFlexibleAppVersion#min_idle_instances}
	MinIdleInstances *float64 `field:"optional" json:"minIdleInstances" yaml:"minIdleInstances"`
	// Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#min_pending_latency GoogleAppEngineFlexibleAppVersion#min_pending_latency}
	MinPendingLatency *string `field:"optional" json:"minPendingLatency" yaml:"minPendingLatency"`
	// Minimum number of running instances that should be maintained for this version. Default: 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#min_total_instances GoogleAppEngineFlexibleAppVersion#min_total_instances}
	MinTotalInstances *float64 `field:"optional" json:"minTotalInstances" yaml:"minTotalInstances"`
	// network_utilization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#network_utilization GoogleAppEngineFlexibleAppVersion#network_utilization}
	NetworkUtilization *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization `field:"optional" json:"networkUtilization" yaml:"networkUtilization"`
	// request_utilization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#request_utilization GoogleAppEngineFlexibleAppVersion#request_utilization}
	RequestUtilization *GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilization `field:"optional" json:"requestUtilization" yaml:"requestUtilization"`
}

