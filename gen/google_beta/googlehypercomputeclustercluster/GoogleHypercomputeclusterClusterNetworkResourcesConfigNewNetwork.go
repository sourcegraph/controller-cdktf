package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterNetworkResourcesConfigNewNetwork struct {
	// Name of the network to create, in the format 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#network GoogleHypercomputeclusterCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Description of the network. Maximum of 2048 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#description GoogleHypercomputeclusterCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

