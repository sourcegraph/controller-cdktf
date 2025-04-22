package googlecontainercluster


type GoogleContainerClusterMonitoringConfigManagedPrometheusAutoMonitoringConfig struct {
	// The scope of auto-monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#scope GoogleContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

