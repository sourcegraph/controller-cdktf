package googlecontainercluster


type GoogleContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig struct {
	// Whether or not the advanced datapath metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#enable_metrics GoogleContainerCluster#enable_metrics}
	EnableMetrics interface{} `field:"required" json:"enableMetrics" yaml:"enableMetrics"`
	// Whether or not Relay is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_cluster#enable_relay GoogleContainerCluster#enable_relay}
	EnableRelay interface{} `field:"required" json:"enableRelay" yaml:"enableRelay"`
}

