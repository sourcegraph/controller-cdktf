package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterNetworkResourcesConfig struct {
	// existing_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#existing_network GoogleHypercomputeclusterCluster#existing_network}
	ExistingNetwork *GoogleHypercomputeclusterClusterNetworkResourcesConfigExistingNetwork `field:"optional" json:"existingNetwork" yaml:"existingNetwork"`
	// new_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_network GoogleHypercomputeclusterCluster#new_network}
	NewNetwork *GoogleHypercomputeclusterClusterNetworkResourcesConfigNewNetwork `field:"optional" json:"newNetwork" yaml:"newNetwork"`
}

