package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterComputeResourcesConfig struct {
	// new_flex_start_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_flex_start_instances GoogleHypercomputeclusterCluster#new_flex_start_instances}
	NewFlexStartInstances *GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances `field:"optional" json:"newFlexStartInstances" yaml:"newFlexStartInstances"`
	// new_on_demand_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_on_demand_instances GoogleHypercomputeclusterCluster#new_on_demand_instances}
	NewOnDemandInstances *GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances `field:"optional" json:"newOnDemandInstances" yaml:"newOnDemandInstances"`
	// new_reserved_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_reserved_instances GoogleHypercomputeclusterCluster#new_reserved_instances}
	NewReservedInstances *GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances `field:"optional" json:"newReservedInstances" yaml:"newReservedInstances"`
	// new_spot_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_spot_instances GoogleHypercomputeclusterCluster#new_spot_instances}
	NewSpotInstances *GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstances `field:"optional" json:"newSpotInstances" yaml:"newSpotInstances"`
}

