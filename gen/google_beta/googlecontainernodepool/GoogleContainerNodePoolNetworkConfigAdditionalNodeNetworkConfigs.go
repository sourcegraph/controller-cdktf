package googlecontainernodepool


type GoogleContainerNodePoolNetworkConfigAdditionalNodeNetworkConfigs struct {
	// Name of the VPC where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#network GoogleContainerNodePool#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Name of the subnetwork where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#subnetwork GoogleContainerNodePool#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

