package googlecontainercluster


type GoogleContainerClusterNodePoolNetworkConfigAdditionalNodeNetworkConfigs struct {
	// Name of the VPC where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#network GoogleContainerCluster#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Name of the subnetwork where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#subnetwork GoogleContainerCluster#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

