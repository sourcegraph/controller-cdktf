package filestoreinstance


type FilestoreInstanceNetworksPscConfig struct {
	// Consumer service project in which the Private Service Connect endpoint would be set up.
	//
	// This is optional, and only relevant in case the network
	// is a shared VPC. If this is not specified, the endpoint would be set up
	// in the VPC host project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/filestore_instance#endpoint_project FilestoreInstance#endpoint_project}
	EndpointProject *string `field:"optional" json:"endpointProject" yaml:"endpointProject"`
}

