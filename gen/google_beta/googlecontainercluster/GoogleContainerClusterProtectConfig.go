package googlecontainercluster


type GoogleContainerClusterProtectConfig struct {
	// workload_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#workload_config GoogleContainerCluster#workload_config}
	WorkloadConfig *GoogleContainerClusterProtectConfigWorkloadConfig `field:"optional" json:"workloadConfig" yaml:"workloadConfig"`
	// Sets which mode to use for Protect workload vulnerability scanning feature. Accepted values are DISABLED, BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#workload_vulnerability_mode GoogleContainerCluster#workload_vulnerability_mode}
	WorkloadVulnerabilityMode *string `field:"optional" json:"workloadVulnerabilityMode" yaml:"workloadVulnerabilityMode"`
}

