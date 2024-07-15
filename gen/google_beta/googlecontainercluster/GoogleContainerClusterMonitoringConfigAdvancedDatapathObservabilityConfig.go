package googlecontainercluster


type GoogleContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig struct {
	// Whether or not the advanced datapath metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enable_metrics GoogleContainerCluster#enable_metrics}
	EnableMetrics interface{} `field:"required" json:"enableMetrics" yaml:"enableMetrics"`
	// Whether or not Relay is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enable_relay GoogleContainerCluster#enable_relay}
	EnableRelay interface{} `field:"optional" json:"enableRelay" yaml:"enableRelay"`
	// Mode used to make Relay available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#relay_mode GoogleContainerCluster#relay_mode}
	RelayMode *string `field:"optional" json:"relayMode" yaml:"relayMode"`
}

