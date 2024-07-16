package apphubworkload


type ApphubWorkloadAttributesBusinessOwners struct {
	// Email address of the contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/apphub_workload#email ApphubWorkload#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Contact's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/apphub_workload#display_name ApphubWorkload#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

