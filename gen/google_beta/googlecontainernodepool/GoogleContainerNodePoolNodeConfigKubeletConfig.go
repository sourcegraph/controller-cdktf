package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigKubeletConfig struct {
	// Control the CPU management policy on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#cpu_manager_policy GoogleContainerNodePool#cpu_manager_policy}
	CpuManagerPolicy *string `field:"required" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#cpu_cfs_quota GoogleContainerNodePool#cpu_cfs_quota}
	CpuCfsQuota interface{} `field:"optional" json:"cpuCfsQuota" yaml:"cpuCfsQuota"`
	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#cpu_cfs_quota_period GoogleContainerNodePool#cpu_cfs_quota_period}
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Controls the maximum number of processes allowed to run in a pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#pod_pids_limit GoogleContainerNodePool#pod_pids_limit}
	PodPidsLimit *float64 `field:"optional" json:"podPidsLimit" yaml:"podPidsLimit"`
}

