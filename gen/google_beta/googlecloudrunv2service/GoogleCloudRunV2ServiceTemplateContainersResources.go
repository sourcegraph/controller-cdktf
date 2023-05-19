package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersResources struct {
	// Determines whether CPU should be throttled or not outside of requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_service#cpu_idle GoogleCloudRunV2Service#cpu_idle}
	CpuIdle interface{} `field:"optional" json:"cpuIdle" yaml:"cpuIdle"`
	// Only memory and CPU are supported.
	//
	// Note: The only supported values for CPU are '1', '2', '4', and '8'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_service#limits GoogleCloudRunV2Service#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Determines whether CPU should be boosted on startup of a new container instance above the requested CPU threshold, this can help reduce cold-start latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_service#startup_cpu_boost GoogleCloudRunV2Service#startup_cpu_boost}
	StartupCpuBoost interface{} `field:"optional" json:"startupCpuBoost" yaml:"startupCpuBoost"`
}

