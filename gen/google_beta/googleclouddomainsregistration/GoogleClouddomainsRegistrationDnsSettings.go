package googleclouddomainsregistration


type GoogleClouddomainsRegistrationDnsSettings struct {
	// custom_dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#custom_dns GoogleClouddomainsRegistration#custom_dns}
	CustomDns *GoogleClouddomainsRegistrationDnsSettingsCustomDns `field:"optional" json:"customDns" yaml:"customDns"`
	// glue_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#glue_records GoogleClouddomainsRegistration#glue_records}
	GlueRecords interface{} `field:"optional" json:"glueRecords" yaml:"glueRecords"`
}

