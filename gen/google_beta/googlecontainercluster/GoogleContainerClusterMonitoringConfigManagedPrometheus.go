package googlecontainercluster


type GoogleContainerClusterMonitoringConfigManagedPrometheus struct {
	// Whether or not the managed collection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

