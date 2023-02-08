package googlecontainercluster


type GoogleContainerClusterMonitoringConfig struct {
	// GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enable_components GoogleContainerCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
	// managed_prometheus block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#managed_prometheus GoogleContainerCluster#managed_prometheus}
	ManagedPrometheus *GoogleContainerClusterMonitoringConfigManagedPrometheus `field:"optional" json:"managedPrometheus" yaml:"managedPrometheus"`
}

