package googlecontainercluster


type GoogleContainerClusterAddonsConfigRayOperatorConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// ray_cluster_logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#ray_cluster_logging_config GoogleContainerCluster#ray_cluster_logging_config}
	RayClusterLoggingConfig *GoogleContainerClusterAddonsConfigRayOperatorConfigRayClusterLoggingConfig `field:"optional" json:"rayClusterLoggingConfig" yaml:"rayClusterLoggingConfig"`
	// ray_cluster_monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#ray_cluster_monitoring_config GoogleContainerCluster#ray_cluster_monitoring_config}
	RayClusterMonitoringConfig *GoogleContainerClusterAddonsConfigRayOperatorConfigRayClusterMonitoringConfig `field:"optional" json:"rayClusterMonitoringConfig" yaml:"rayClusterMonitoringConfig"`
}

