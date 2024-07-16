package googleclouddomainsregistration


type GoogleClouddomainsRegistrationContactSettingsRegistrantContact struct {
	// Required. Email address of the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#email GoogleClouddomainsRegistration#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Required. Phone number of the contact in international format. For example, "+1-800-555-0123".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#phone_number GoogleClouddomainsRegistration#phone_number}
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// postal_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#postal_address GoogleClouddomainsRegistration#postal_address}
	PostalAddress *GoogleClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress `field:"required" json:"postalAddress" yaml:"postalAddress"`
	// Fax number of the contact in international format. For example, "+1-800-555-0123".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#fax_number GoogleClouddomainsRegistration#fax_number}
	FaxNumber *string `field:"optional" json:"faxNumber" yaml:"faxNumber"`
}

