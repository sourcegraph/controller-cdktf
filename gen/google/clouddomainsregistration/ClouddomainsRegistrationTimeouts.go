package clouddomainsregistration


type ClouddomainsRegistrationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/clouddomains_registration#create ClouddomainsRegistration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/clouddomains_registration#delete ClouddomainsRegistration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/clouddomains_registration#update ClouddomainsRegistration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

