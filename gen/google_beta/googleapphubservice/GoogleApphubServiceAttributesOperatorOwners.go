package googleapphubservice


type GoogleApphubServiceAttributesOperatorOwners struct {
	// Required. Email address of the contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#email GoogleApphubService#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Contact's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#display_name GoogleApphubService#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

