package googleclouddomainsregistration


type GoogleClouddomainsRegistrationContactSettings struct {
	// admin_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#admin_contact GoogleClouddomainsRegistration#admin_contact}
	AdminContact *GoogleClouddomainsRegistrationContactSettingsAdminContact `field:"required" json:"adminContact" yaml:"adminContact"`
	// Required. Privacy setting for the contacts associated with the Registration. Values are PUBLIC_CONTACT_DATA, PRIVATE_CONTACT_DATA, and REDACTED_CONTACT_DATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#privacy GoogleClouddomainsRegistration#privacy}
	Privacy *string `field:"required" json:"privacy" yaml:"privacy"`
	// registrant_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#registrant_contact GoogleClouddomainsRegistration#registrant_contact}
	RegistrantContact *GoogleClouddomainsRegistrationContactSettingsRegistrantContact `field:"required" json:"registrantContact" yaml:"registrantContact"`
	// technical_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#technical_contact GoogleClouddomainsRegistration#technical_contact}
	TechnicalContact *GoogleClouddomainsRegistrationContactSettingsTechnicalContact `field:"required" json:"technicalContact" yaml:"technicalContact"`
}

