package hypercomputeclustercluster


type HypercomputeclusterClusterComputeResourcesConfig struct {
	// new_flex_start_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_flex_start_instances HypercomputeclusterCluster#new_flex_start_instances}
	NewFlexStartInstances *HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances `field:"optional" json:"newFlexStartInstances" yaml:"newFlexStartInstances"`
	// new_on_demand_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_on_demand_instances HypercomputeclusterCluster#new_on_demand_instances}
	NewOnDemandInstances *HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances `field:"optional" json:"newOnDemandInstances" yaml:"newOnDemandInstances"`
	// new_reserved_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_reserved_instances HypercomputeclusterCluster#new_reserved_instances}
	NewReservedInstances *HypercomputeclusterClusterComputeResourcesConfigNewReservedInstances `field:"optional" json:"newReservedInstances" yaml:"newReservedInstances"`
	// new_spot_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_spot_instances HypercomputeclusterCluster#new_spot_instances}
	NewSpotInstances *HypercomputeclusterClusterComputeResourcesConfigNewSpotInstances `field:"optional" json:"newSpotInstances" yaml:"newSpotInstances"`
}

