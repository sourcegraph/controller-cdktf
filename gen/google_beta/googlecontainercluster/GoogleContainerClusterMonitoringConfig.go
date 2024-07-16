package googlecontainercluster


type GoogleContainerClusterMonitoringConfig struct {
	// advanced_datapath_observability_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#advanced_datapath_observability_config GoogleContainerCluster#advanced_datapath_observability_config}
	AdvancedDatapathObservabilityConfig *GoogleContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig `field:"optional" json:"advancedDatapathObservabilityConfig" yaml:"advancedDatapathObservabilityConfig"`
	// GKE components exposing metrics.
	//
	// Valid values include SYSTEM_COMPONENTS, APISERVER, SCHEDULER, CONTROLLER_MANAGER, STORAGE, HPA, POD, DAEMONSET, DEPLOYMENT, STATEFULSET, WORKLOADS, KUBELET, CADVISOR and DCGM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enable_components GoogleContainerCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
	// managed_prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#managed_prometheus GoogleContainerCluster#managed_prometheus}
	ManagedPrometheus *GoogleContainerClusterMonitoringConfigManagedPrometheus `field:"optional" json:"managedPrometheus" yaml:"managedPrometheus"`
}

