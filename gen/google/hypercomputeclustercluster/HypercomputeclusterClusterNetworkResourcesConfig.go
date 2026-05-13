package hypercomputeclustercluster


type HypercomputeclusterClusterNetworkResourcesConfig struct {
	// existing_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#existing_network HypercomputeclusterCluster#existing_network}
	ExistingNetwork *HypercomputeclusterClusterNetworkResourcesConfigExistingNetwork `field:"optional" json:"existingNetwork" yaml:"existingNetwork"`
	// new_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_network HypercomputeclusterCluster#new_network}
	NewNetwork *HypercomputeclusterClusterNetworkResourcesConfigNewNetwork `field:"optional" json:"newNetwork" yaml:"newNetwork"`
}

