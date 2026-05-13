package cloudfunctions2function


type Cloudfunctions2FunctionServiceConfigDirectVpcNetworkInterface struct {
	// The name of the VPC network to which the function will be connected.
	//
	// Specify either a VPC network or a subnet, or both. If you specify only a network, the subnet uses the same name as the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudfunctions2_function#network Cloudfunctions2Function#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the VPC subnetwork that the Cloud Function resource will get IPs from.
	//
	// Specify either a VPC network or a subnet, or both. If both network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If subnetwork is not specified, the subnetwork with the same name with the network will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudfunctions2_function#subnetwork Cloudfunctions2Function#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// Network tags applied to this Cloud Function resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudfunctions2_function#tags Cloudfunctions2Function#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

