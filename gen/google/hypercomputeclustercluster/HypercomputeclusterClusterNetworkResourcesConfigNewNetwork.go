package hypercomputeclustercluster


type HypercomputeclusterClusterNetworkResourcesConfigNewNetwork struct {
	// Name of the network to create, in the format 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#network HypercomputeclusterCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Description of the network. Maximum of 2048 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#description HypercomputeclusterCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

