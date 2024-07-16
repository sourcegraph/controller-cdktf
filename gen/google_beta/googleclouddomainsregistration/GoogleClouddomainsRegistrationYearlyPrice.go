package googleclouddomainsregistration


type GoogleClouddomainsRegistrationYearlyPrice struct {
	// The three-letter currency code defined in ISO 4217.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#currency_code GoogleClouddomainsRegistration#currency_code}
	CurrencyCode *string `field:"optional" json:"currencyCode" yaml:"currencyCode"`
	// The whole units of the amount. For example if currencyCode is "USD", then 1 unit is one US dollar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#units GoogleClouddomainsRegistration#units}
	Units *string `field:"optional" json:"units" yaml:"units"`
}

