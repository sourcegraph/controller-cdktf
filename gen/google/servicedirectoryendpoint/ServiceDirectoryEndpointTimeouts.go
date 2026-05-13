package servicedirectoryendpoint


type ServiceDirectoryEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_endpoint#create ServiceDirectoryEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_endpoint#delete ServiceDirectoryEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_endpoint#update ServiceDirectoryEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

