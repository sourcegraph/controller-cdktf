package googleapphubapplication


type GoogleApphubApplicationAttributesBusinessOwners struct {
	// Required. Email address of the contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_application#email GoogleApphubApplication#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Optional. Contact's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_application#display_name GoogleApphubApplication#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

