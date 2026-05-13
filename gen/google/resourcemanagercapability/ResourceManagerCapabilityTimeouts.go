package resourcemanagercapability


type ResourceManagerCapabilityTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/resource_manager_capability#create ResourceManagerCapability#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/resource_manager_capability#delete ResourceManagerCapability#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/resource_manager_capability#update ResourceManagerCapability#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

