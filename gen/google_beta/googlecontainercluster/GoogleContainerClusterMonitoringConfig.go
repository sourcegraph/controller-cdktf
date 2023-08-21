package googlecontainercluster


type GoogleContainerClusterMonitoringConfig struct {
	// advanced_datapath_observability_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_container_cluster#advanced_datapath_observability_config GoogleContainerCluster#advanced_datapath_observability_config}
	AdvancedDatapathObservabilityConfig interface{} `field:"optional" json:"advancedDatapathObservabilityConfig" yaml:"advancedDatapathObservabilityConfig"`
	// GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_container_cluster#enable_components GoogleContainerCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
	// managed_prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_container_cluster#managed_prometheus GoogleContainerCluster#managed_prometheus}
	ManagedPrometheus *GoogleContainerClusterMonitoringConfigManagedPrometheus `field:"optional" json:"managedPrometheus" yaml:"managedPrometheus"`
}

