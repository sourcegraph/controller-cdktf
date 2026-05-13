package hypercomputeclustercluster


type HypercomputeclusterClusterNetworkResourcesConfigExistingNetwork struct {
	// Name of the network to import, in the format 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#network HypercomputeclusterCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Particular subnetwork to use, in the format 'projects/{project}/regions/{region}/subnetworks/{subnetwork}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#subnetwork HypercomputeclusterCluster#subnetwork}
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
}

